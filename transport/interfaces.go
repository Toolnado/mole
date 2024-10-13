package transport

import (
	"io"
	"mole/model"
)

type Transport interface {
	ListenAndServe() error
}

type Decoder interface {
	Decode(name string, r io.Reader) (model.RPC, error)
}
