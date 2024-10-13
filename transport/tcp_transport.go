package transport

import (
	"mole/logger"
	"mole/peer"
	"net"
	"sync"
)

type TCPTransport struct {
	Address string
	Options Options
	mu      sync.RWMutex
	peers   map[net.Addr]peer.Peer
}

type Options struct {
	Listener  net.Listener
	Logger    logger.Logger
	Decoder   Decoder
	Handshake func(peer.Peer) error
	OnPeer    func(peer.Peer) error
}

func NewTCPTransport(opts Options) TCPTransport {
	return TCPTransport{
		Address: opts.Listener.Addr().String(),
		Options: opts,

		mu:    sync.RWMutex{},
		peers: make(map[net.Addr]peer.Peer),
	}
}

func (tcpt *TCPTransport) ListenAndServe() error {
	for {
		conn, err := tcpt.Options.Listener.Accept()
		if err != nil {
			tcpt.Options.Logger.Error("accept connection error: %s", err)
			continue
		}
		tcpt.Options.Logger.Info("accept new connection: %s", conn.RemoteAddr().String())

		tcpPeer := peer.NewTCPPeer(conn)

		if err = tcpt.Options.Handshake(&tcpPeer); err != nil {
			tcpt.Options.Logger.Error("handshake connection error: %s", err)
			continue
		}

		if err = tcpt.Options.OnPeer(&tcpPeer); err != nil {
			tcpt.Options.Logger.Error("on peer connection error: %s", err)
			continue
		}

		go tcpt.Handle(&tcpPeer)
	}
}

func (tcpt *TCPTransport) Handle(p peer.Peer) {
	for {
		if msg, err := tcpt.Options.Decoder.Decode(p.Addr(), p.NewReader()); err != nil {
			tcpt.Options.Logger.Error("decode message error: %s", err)
		} else {
			tcpt.Options.Logger.Info("receive message: %s", msg)
		}
	}
}
