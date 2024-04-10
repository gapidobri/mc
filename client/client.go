package client

import (
	"bytes"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"mc/packet"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *packet.Reader
	writer *packet.Writer

	state      packet.State
	privateKey *rsa.PrivateKey

	player *Player
}

func NewClient(conn net.Conn, privateKey *rsa.PrivateKey) *Client {
	return &Client{
		conn:       conn,
		reader:     packet.NewReader(conn),
		writer:     packet.NewWriter(conn),
		state:      packet.StateHandshaking,
		privateKey: privateKey,
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
		if packet.Type(packetId) == packet.Handshake {
			return c.handleHandshake(reader)
		}

	case packet.StateStatus:
		return c.handleStatusState(packetId, reader)

	case packet.StateLogin:
		return c.handleLoginState(packetId, reader)

	case packet.StateConfiguration:
		return c.handleConfigurationState(packetId, reader)
	}

	fmt.Println("unknown packet")
	fmt.Printf("packet id: 0x%s\n", hex.EncodeToString([]byte{byte(packetId)}))
	fmt.Printf("0x%s\n", hex.EncodeToString(data))

	return nil
}

func (c *Client) reply(p packet.Packet) error {
	buf := bytes.NewBuffer(nil)
	w := packet.NewWriter(buf)

	err := w.WriteVarInt(p.PacketId())
	err = packet.Write(w, p)
	err = w.Flush()
	if err != nil {
		return err
	}

	err = c.writer.WriteVarInt(buf.Len())
	if err != nil {
		return err
	}

	_, err = c.writer.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return c.writer.Flush()
}

func (c *Client) handleHandshake(r *packet.Reader) error {
	var handshakeReq packet.HandshakeReq
	err := packet.Read(r, &handshakeReq)
	if err != nil {
		return err
	}

	fmt.Printf("< handshake request: %+v\n", handshakeReq)

	c.state = handshakeReq.NextState

	return nil
}
