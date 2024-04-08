package client

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"mc/packet"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *packet.Reader
	writer *packet.Writer

	state packet.State
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:   conn,
		reader: packet.NewReader(conn),
		writer: packet.NewWriter(conn),
		state:  packet.StateHandshaking,
	}
}

func (c *Client) HandleRequest() {
	defer c.conn.Close()

	fmt.Println("Handling request from", c.conn.RemoteAddr())

	for {
		err := c.readPacket()
		if err != nil {
			fmt.Println("an error occured:", err)
			break
		}
	}
}

func (c *Client) readPacket() error {
	fmt.Println()

	// Read content length
	length, err := c.reader.ReadVarInt()
	if err != nil {
		return err
	}

	// Read the packet contents
	data := make([]byte, length)
	_, err = c.reader.Read(data)
	if err != nil {
		return err
	}

	reader := packet.NewReader(bytes.NewReader(data))

	packetId, err := reader.ReadVarInt()
	if err != nil {
		return err
	}

	switch c.state {
	case packet.StateHandshaking:
		switch packet.Type(packetId) {
		case packet.Handshake:
			return c.handleHandshake(reader)
		}

	case packet.StateStatus:
		switch packet.Type(packetId) {
		case packet.Status:
			return c.handleStatus()
		case packet.Ping:
			return c.handlePing(reader)
		}

	case packet.StateLogin:
		switch packet.Type(packetId) {
		case packet.Login:
			return c.handleLogin(reader)
		}
	}

	fmt.Println("unknown packet")
	fmt.Printf("0x%s\n", hex.EncodeToString(data))
	fmt.Printf("packet id: 0x%s\n", hex.EncodeToString([]byte{byte(packetId)}))

	return nil
}

func (c *Client) reply(packet packet.Packet) error {
	data, err := packet.Encode()
	if err != nil {
		return err
	}

	err = c.writer.WriteVarInt(len(data))
	if err != nil {
		return err
	}

	_, err = c.writer.Write(data)
	if err != nil {
		return err
	}

	return c.writer.Flush()
}

func (c *Client) handleHandshake(r *packet.Reader) error {
	handshakeReq, err := packet.ReadHandshakeReq(r)
	if err != nil {
		return err
	}

	fmt.Printf("handshake request: %+v\n", handshakeReq)

	c.state = handshakeReq.NextState

	return nil
}

func (c *Client) handleStatus() error {
	fmt.Println("status request")
	statusRes := packet.StatusRes{
		Version: packet.Version{
			Name:     "1.20.4",
			Protocol: 762,
		},
		Players: packet.Players{
			Max:    420,
			Online: 69,
			Sample: []packet.Sample{
				{
					Name: "tvoja_mami",
					Id:   "4566e69f-c907-48ee-8d71-d7ba5aa00d20",
				},
			},
		},
		Description: packet.Description{
			Text: "It works!!!",
		},
	}

	fmt.Printf("status reply: %+v\n", statusRes)
	return c.reply(statusRes)
}

func (c *Client) handlePing(r *packet.Reader) error {
	pingReq, err := packet.ReadPingReq(r)
	if err != nil {
		return err
	}

	fmt.Printf("ping request: %+v\n", pingReq)

	pingRes := packet.PingRes{
		Payload: pingReq.Payload,
	}

	fmt.Printf("ping reply: %+v\n", pingRes)

	return c.reply(pingRes)
}

func (c *Client) handleLogin(r *packet.Reader) error {
	loginReq, err := packet.ReadLoginReq(r)
	if err != nil {
		return err
	}

	fmt.Printf("login request: %+v\n", loginReq)

	return nil
}
