package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	transport := transport.NewTCPTransport(
		transport.Context{
			Value:  ctx,
			Cancel: cancel,
		},
		transport.NewComponents(
			listener,
			log,
			decoder,
			transport.NopSecurity{},
			transport.NopAcceptance{},
		),
	)
	go transport.ListenAndServe()
	transport.Wait()
	transport.Close()
}
