package main

import (
	"mole/logger"
	"mole/transport"
	"net"
)

func main() {
	l := logger.New()
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	opts := transport.Options{
		Listener:  li,
		Logger:    l,
		Decoder:   transport.DefautlDecoder,
		Handshake: transport.NOPHandshake,
		OnPeer:    transport.NOPOnPeer,
	}
	transport := transport.NewTCPTransport(opts)
	transport.ListenAndServe()
	select {}
}
