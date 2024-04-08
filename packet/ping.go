package packet

import "bytes"

type PingReq struct {
	Payload int64
}

func ReadPingReq(r *Reader) (*PingReq, error) {
	payload, err := r.ReadInt64()
	if err != nil {
		return nil, err
	}

	pingReq := &PingReq{
		Payload: payload,
	}

	return pingReq, nil
}

type PingRes struct {
	Payload int64
}

func (p PingRes) Encode() ([]byte, error) {
	data := bytes.NewBuffer(nil)
	writer := NewWriter(data)

	err := writer.WriteInt64(p.Payload)
	if err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
