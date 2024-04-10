package client

import (
	"encoding/hex"
	"fmt"
	"mc/packet"
)

func (c *Client) handleConfigurationState(packetId int, r *packet.Reader) error {
	switch packet.Type(packetId) {
	case packet.ConfigClientInfo:
		return c.handleConfigClientInfo(r)
	case packet.ConfigPluginMessage:
		return c.handleConfigPluginMessage(r)
	case packet.FinishConfig:
		return c.handleFinishConfigAck()
	}
	return fmt.Errorf("unknown configuration packet: 0x%s", hex.EncodeToString([]byte{byte(packetId)}))
}

func (c *Client) handleConfigClientInfo(r *packet.Reader) error {
	var clientInfo packet.ClientInfo
	err := packet.Read(r, &clientInfo)
	if err != nil {
		return err
	}

	fmt.Printf("< client info: %+v\n", clientInfo)

	fmt.Println("> finish config request")

	return c.reply(packet.FinishConfigReq{})
}

func (c *Client) handleConfigPluginMessage(r *packet.Reader) error {
	pluginMessage, err := packet.ReadPluginMessage(r)
	if err != nil {
		return err
	}

	fmt.Printf("< plugin message: %+v\n", pluginMessage)

	return nil
}

func (c *Client) handleFinishConfigAck() error {
	fmt.Println("< finish config ack")

	c.state = packet.StatePlay

	return nil
}
