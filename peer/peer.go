package peer

import (
	"io"
)

type Peer interface {
	Addr() string
	Close() error
	Reader() io.Reader
}
