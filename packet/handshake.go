package packet

type HandshakeReq struct {
	ProtocolVersion int
	ServerAddress   string
	ServerPort      uint16
	NextState       State
}

func ReadHandshakeReq(r *Reader) (*HandshakeReq, error) {
	protocolVersion, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	serverAddress, err := r.ReadString()
	if err != nil {
		return nil, err
	}

	serverPort, err := r.ReadUint16()
	if err != nil {
		return nil, err
	}

	nextState, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	handshakeReq := &HandshakeReq{
		ProtocolVersion: protocolVersion,
		ServerAddress:   serverAddress,
		ServerPort:      serverPort,
		NextState:       State(nextState),
	}

	return handshakeReq, nil
}
