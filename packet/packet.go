package packet

type Packet interface {
	Encode() ([]byte, error)
}
