package packet

type FinishConfigReq struct{}

func (FinishConfigReq) PacketId() int {
	return int(FinishConfig)
}
