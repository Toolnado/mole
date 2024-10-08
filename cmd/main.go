package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	littleMole1 := NewMole("localhost:8081")
	littleMole2 := NewMole("localhost:8082")

	go littleMole2.Listen()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(dir, "from", "test.txt")
	littleMole1.SendFile("localhost:8082", filename)

	select {}
}

type Mole struct {
	Address string
}

func NewMole(address string) *Mole {
	return &Mole{Address: address}
}

func (m *Mole) Listen() {
	protoError := errors.New("mole.Listen error")
	listener, err := net.Listen("tcp", m.Address)
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(errors.Join(protoError, err))
			continue
		}
		go m.handleConn(conn)
	}
}

func (m *Mole) handleConn(conn net.Conn) {
	protoError := errors.New("mole.handleConn error")
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(errors.Join(protoError, err))
	}
	filename := string(buf[:n])
	fmt.Println(filename)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(errors.Join(protoError, err))
	}
	file, err := os.Create(filepath.Join(dir, "to", filename))
	if err != nil {
		log.Println(errors.Join(protoError, err))
	}
	defer file.Close()
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	_, err = file.Write([]byte("new "))
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	_, err = file.Write(buf[:n])
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	log.Printf("file %s received successfully.\n", filename)
}

func (p *Mole) SendFile(address string, filepath string) {
	protoError := errors.New("mole.SendFile error")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	defer conn.Close()
	filename, err := os.Stat(filepath)
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	conn.Write([]byte(filename.Name()))
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	_, err = conn.Write(buf[:n])
	if err != nil {
		log.Println(errors.Join(protoError, err))
		return
	}
	log.Printf("File %s sent to %s.\n", filepath, address)
}
