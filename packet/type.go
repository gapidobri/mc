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
	Login        Type = 0x00
	Encryption   Type = 0x01
	LoginSuccess Type = 0x02
	LoginAck     Type = 0x03
)

// Configuration

const (
	ConfigClientInfo    Type = 0x00
	ConfigPluginMessage Type = 0x01
	FinishConfig        Type = 0x02
)
