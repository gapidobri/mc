package packet

type UUID [16]byte

type LoginReq struct {
	Name       string
	PlayerUUID UUID
}

func ReadLoginReq(r *Reader) (*LoginReq, error) {
	name, err := r.ReadString()
	if err != nil {
		return nil, err
	}

	playerUUID, err := r.ReadUUID()
	if err != nil {
		return nil, err
	}

	loginReq := &LoginReq{
		Name:       name,
		PlayerUUID: playerUUID,
	}

	return loginReq, nil
}
