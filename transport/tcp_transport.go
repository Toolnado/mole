package transport

import (
	"net"
	"sync"

	"github.com/Toolnado/mole/interfaces"
	"github.com/Toolnado/mole/peer"
)

type TCPTransport struct {
	Address    string
	Components Components
	mu         sync.RWMutex
	peers      map[net.Addr]interfaces.Peer
}

func NewTCPTransport(opts Components) TCPTransport {
	return TCPTransport{
		Address:    opts.Listener.Addr().String(),
		Components: opts,

		mu:    sync.RWMutex{},
		peers: make(map[net.Addr]interfaces.Peer),
	}
}

func (tcpt *TCPTransport) ListenAndServe() error {
	for {
		conn, err := tcpt.Components.Listener.Accept()
		if err != nil {
			tcpt.Components.Logger.Error("accept connection error: %s", err)
			continue
		}
		tcpt.Components.Logger.Info("accept new connection: %s", conn.RemoteAddr().String())

		tcpPeer := peer.NewTCPPeer(conn)

		if err = tcpt.Components.Security.Handshake(&tcpPeer); err != nil {
			tcpt.Components.Logger.Error("handshake connection error: %s", err)
			continue
		}

		if err = tcpt.Components.Acceptance.OnPeer(&tcpPeer); err != nil {
			tcpt.Components.Logger.Error("on peer connection error: %s", err)
			continue
		}

		go tcpt.handle(&tcpPeer)
	}
}

func (tcpt *TCPTransport) handle(p interfaces.Peer) {
	for {
		if msg, err := tcpt.Components.Decoder.Decode(p.Addr(), p.Reader()); err != nil {
			tcpt.Components.Logger.Error("decode message error: %s", err)
		} else {
			tcpt.Components.Logger.Info("receive message: %s", msg)
		}
	}
}
