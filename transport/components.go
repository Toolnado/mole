package transport

import (
	"net"

	"github.com/Toolnado/mole/interfaces"
)

type Components struct {
	Listener   net.Listener
	Logger     interfaces.Logger
	Decoder    interfaces.Decoder
	Security   interfaces.Security
	Acceptance interfaces.Acceptance
}

func NewComponents(
	li net.Listener,
	log interfaces.Logger,
	dec interfaces.Decoder,
	sec interfaces.Security,
	acc interfaces.Acceptance,
) Components {
	return Components{
		Listener:   li,
		Logger:     log,
		Decoder:    dec,
		Security:   sec,
		Acceptance: acc,
	}
}
