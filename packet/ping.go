package packet

type PingReq struct {
	Payload int64
}

type PingRes struct {
	Payload int64
}

func (p PingRes) PacketId() int {
	return int(Ping)
}
