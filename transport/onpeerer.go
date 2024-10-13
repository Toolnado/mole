package transport

import "mole/peer"

func NOPOnPeer(peer.Peer) error {
	return nil
}
