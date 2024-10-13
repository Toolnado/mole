package transport

import (
	"io"
	"mole/model"
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
	msg := model.RPC{PeerAddress: name}
	if _, err := r.Read(buf); err != nil {
		if err != io.EOF {
			return model.RPC{}, err
		}
	} else {
		msg.Payload = buf
	}
	return msg, nil
}
