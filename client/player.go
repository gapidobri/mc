package client

import "mc/packet"

type Player struct {
	Username string
	UUID     packet.UUID
}
