package mole

import (
	"bufio"
	"context"
	"errors"
	"io"
	"mole/pkg/logger"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

type Mole struct {
	ctx     context.Context
	cancel  context.CancelFunc
	address string
	logger  *logger.Logger
}

func NewMole(address string) *Mole {
	ctx, cancel := context.WithCancel(context.Background())
	logger := logger.New()
	return &Mole{
		ctx:     ctx,
		cancel:  cancel,
		address: address,
		logger:  logger,
	}
}

func (m *Mole) Wait() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}

func (m *Mole) Stop() {
	m.cancel()
}

func (m *Mole) Listen() error {
	protoError := errors.New("mole.Listen error")
	listener, err := net.Listen("tcp", m.address)
	if err != nil {
		return errors.Join(protoError, err)
	}
	tunnel := make(chan net.Conn)
	go m.handle(tunnel)
	for {
		select {
		case <-m.ctx.Done():
			close(tunnel)
			return nil
		default:
			conn, err := listener.Accept()
			if err != nil {
				m.logger.Error("%s", errors.Join(protoError, err))
				continue
			}
			tunnel <- conn
		}
	}
}

func (m *Mole) handle(connChan chan net.Conn) {
	for {
		select {
		case <-m.ctx.Done():
			return
		case conn := <-connChan:
			go m.handleConn(conn)
		}
	}
}

func (m *Mole) handleConn(conn net.Conn) {
	defer conn.Close()
	m.logger.Info("%s", "handle new conn")
	protoError := errors.New("mole.handleConn error")

	reader := bufio.NewReader(conn)
	filename, err := reader.ReadString('\n')
	if err != nil {
		m.logger.Error("%s", errors.Join(protoError, err))
		return
	}
	filename = filename[:len(filename)-1]
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		m.logger.Error("%s", errors.Join(protoError, err))
		return
	}
	file, err := os.Create(filepath.Join(dir, "assets", filename))
	if err != nil {
		m.logger.Error("%s", errors.Join(protoError, err))
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			m.logger.Error("%s", errors.Join(protoError, err))
			return
		}
		_, err = file.Write(buf[:n])
		if err != nil {
			m.logger.Error("%s", errors.Join(protoError, err))
			return
		}
	}

	m.logger.Info("file %s received successfully.", filename)
}

func (m *Mole) SendFile(address string, filepath string) error {
	protoError := errors.New("mole.SendFile error")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Join(protoError, err)
	}
	defer conn.Close()
	filename, err := os.Stat(filepath)
	if err != nil {
		return errors.Join(protoError, err)
	}
	bytesFilename := append([]byte(filename.Name()), '\n')
	conn.Write(bytesFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return errors.Join(protoError, err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return errors.Join(protoError, err)
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			return errors.Join(protoError, err)
		}
	}
	m.logger.Info("File %s sent to %s.", filepath, address)
	return nil
}
