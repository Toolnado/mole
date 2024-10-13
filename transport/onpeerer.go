package transport

import "github.com/Toolnado/mole/peer"

func NOPOnPeer(peer.Peer) error {
	return nil
}
