package transport

import (
	"io"

	"github.com/Toolnado/mole/model"
)

type DefautlDecoder struct {
	bufSize int
}

func NewDefautlDecoder(bs int) DefautlDecoder {
	return DefautlDecoder{
		bufSize: bs,
	}
}

func (dd DefautlDecoder) Decode(name string, r io.Reader) (model.RPC, error) {
	buf := make([]byte, dd.bufSize)
	if _, err := r.Read(buf); err != nil {
		if err != io.EOF {
			return model.RPC{}, err
		}
	}
	return model.RPC{
		PeerAddress: name,
		Payload:     buf,
	}, nil
}
