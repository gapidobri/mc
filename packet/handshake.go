package packet

type HandshakeReq struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      uint16
	NextState       State
}
