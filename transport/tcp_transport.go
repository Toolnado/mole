package transport

import (
	"context"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Toolnado/mole/interfaces"
	"github.com/Toolnado/mole/peer"
)

type TCPTransport struct {
	Address    string
	Components Components

	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.RWMutex
	peers  map[net.Addr]interfaces.Peer
}

type Context struct {
	Value  context.Context
	Cancel context.CancelFunc
}

func NewTCPTransport(ctx Context, opts Components) TCPTransport {
	return TCPTransport{
		Address:    opts.Listener.Addr().String(),
		Components: opts,

		ctx:    ctx.Value,
		cancel: ctx.Cancel,
		mu:     sync.RWMutex{},
		peers:  make(map[net.Addr]interfaces.Peer),
	}
}

func (tcpt *TCPTransport) handleContextError() (bool, error) {
	if err := tcpt.ctx.Err(); err != nil {
		if err == context.Canceled {
			return true, nil
		}
		return true, err
	}
	return false, nil
}

func (tcpt *TCPTransport) ListenAndServe() error {
	for {
		if close, err := tcpt.handleContextError(); close {
			return err
		}

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
		if close, _ := tcpt.handleContextError(); close {
			return
		}
		if msg, err := tcpt.Components.Decoder.Decode(p.Addr(), p.Reader()); err != nil {
			tcpt.Components.Logger.Error("decode message error: %s", err)
		} else {
			tcpt.Components.Logger.Info("receive message: %s", msg)
		}
	}
}

func (tcpt *TCPTransport) Wait() {
	closer := make(chan os.Signal, 1)
	signal.Notify(closer, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-closer
}

func (tcpt *TCPTransport) Close() {
	tcpt.Components.Logger.Info("closing transport with address: %s", tcpt.Address)
	tcpt.cancel()
	time.Sleep(1 * time.Second)
	tcpt.Components.Logger.Info("transport with address %s closed", tcpt.Address)
}
