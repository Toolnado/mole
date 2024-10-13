package peer

import (
	"io"
	"net"
)

type TCPPeer struct {
	addr string
	conn net.Conn
}

func NewTCPPeer(conn net.Conn) TCPPeer {
	return TCPPeer{
		addr: conn.RemoteAddr().String(),
		conn: conn,
	}
}

func (tcpp *TCPPeer) Close() error {
	return tcpp.conn.Close()
}

func (tcpp *TCPPeer) Addr() string {
	return tcpp.addr
}

func (tcpp TCPPeer) NewReader() io.Reader {
	return tcpp.conn
}
