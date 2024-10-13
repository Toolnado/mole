package transport

import (
	"io"
	"mole/model"
)

func DefautlDecoder(name string, r io.Reader) (model.RPC, error) {
	buf := make([]byte, 1024)
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
