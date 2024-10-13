package transport

import "github.com/Toolnado/mole/interfaces"

type NopAcceptance struct{}

func (acceptance NopAcceptance) OnPeer(interfaces.Peer) error {
	return nil
}
