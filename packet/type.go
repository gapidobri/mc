package packet

type Type int

// Handshaking

const (
	Handshake Type = 0x00
)

// Status

const (
	Status Type = 0x00
	Ping   Type = 0x01
)

// Login

const (
	Login Type = 0x00
)
