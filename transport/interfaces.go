package transport

import (
	"io"

	"github.com/Toolnado/mole/model"
)

type Transport interface {
	ListenAndServe() error
}

type Decoder interface {
	Decode(name string, r io.Reader) (model.RPC, error)
}
