package client

import (
	"encoding/hex"
	"fmt"
	"mc/packet"
)

func (c *Client) handleStatusState(packetId int, r *packet.Reader) error {
	switch packet.Type(packetId) {
	case packet.Status:
		return c.handleStatus()
	case packet.Ping:
		return c.handlePing(r)
	}
	return fmt.Errorf("unknown status packet: 0x%s", hex.EncodeToString([]byte{byte(packetId)}))
}

func (c *Client) handleStatus() error {
	fmt.Println("< status request")

	statusRes := packet.StatusRes{
		Version: packet.Version{
			Name:     "1.20.4",
			Protocol: 762,
		},
		Players: packet.Players{
			Max:    10,
			Online: 1,
			Sample: []packet.Sample{
				{
					Name: "player",
					Id:   "4566e69f-c907-48ee-8d71-d7ba5aa00d20",
				},
			},
		},
		Description: packet.Description{
			Text: "Minecraft server in go",
		},
	}

	fmt.Printf("> status response: %+v\n", statusRes)

	return c.reply(statusRes)
}

func (c *Client) handlePing(r *packet.Reader) error {
	var pingReq packet.PingReq
	err := packet.Read(r, &pingReq)
	if err != nil {
		return err
	}

	fmt.Printf("< ping request: %+v\n", pingReq)

	pingRes := packet.PingRes{
		Payload: pingReq.Payload,
	}

	fmt.Printf("> ping response: %+v\n", pingRes)

	return c.reply(pingRes)
}
