package transport

import (
	"github.com/Toolnado/mole/interfaces"
)

type NopSecurity struct{}

func (security NopSecurity) Handshake(interfaces.Peer) error {
	return nil
}
