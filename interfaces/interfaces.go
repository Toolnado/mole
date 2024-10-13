package interfaces

import (
	"io"

	"github.com/Toolnado/mole/model"
)

type Peer interface {
	Addr() string
	Close() error
	Reader() io.Reader
}

type Transport interface {
	ListenAndServe() error
}

type Decoder interface {
	Decode(name string, r io.Reader) (model.RPC, error)
}

type Security interface {
	Handshake(Peer) error
}

type Logger interface {
	Error(v ...any)
	Fatal(v ...any)
	Info(v ...any)
}

type Acceptance interface {
	OnPeer(Peer) error
}
