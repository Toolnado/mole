package transport

import "github.com/Toolnado/mole/peer"

func NOPHandshake(peer.Peer) error {
	return nil
}
