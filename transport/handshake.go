package transport

import "mole/peer"

func NOPHandshake(peer.Peer) error {
	return nil
}
