package main

import (
	"net"

	"github.com/Toolnado/mole/logger"
	"github.com/Toolnado/mole/transport"
)

func main() {
	log := logger.New()
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	decoder := transport.NewDefautlDecoder(1024)
	transport := transport.NewTCPTransport(
		transport.NewComponents(
			listener,
			log,
			decoder,
			transport.NopSecurity{},
			transport.NopAcceptance{},
		),
	)
	transport.ListenAndServe()
	select {}
}
