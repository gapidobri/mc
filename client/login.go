package client

import (
	"encoding/hex"
	"fmt"
	"mc/packet"
)

func (c *Client) handleLoginState(packetId int, r *packet.Reader) error {
	switch packet.Type(packetId) {
	case packet.Login:
		return c.handleLogin(r)
	case packet.Encryption:
		return c.handleEncryption(r)
	case packet.LoginAck:
		return c.handleLoginAck()
	}
	return fmt.Errorf("unknown login packet: 0x%s", hex.EncodeToString([]byte{byte(packetId)}))
}

func (c *Client) handleLogin(r *packet.Reader) error {
	loginReq, err := packet.ReadLoginReq(r)
	if err != nil {
		return err
	}

	fmt.Printf("< login request: %+v\n", loginReq)

	c.player = &Player{
		Username: loginReq.Name,
		UUID:     loginReq.PlayerUUID,
	}

	//publicKey, err := x509.MarshalPKIXPublicKey(&c.privateKey.PublicKey)
	//if err != nil {
	//	return err
	//}
	//
	//verifyToken := make([]byte, 4)
	//_, err = rand.Read(verifyToken)
	//if err != nil {
	//	return err
	//}

	//encryptionReq := packet.EncryptionReq{
	//	PublicKey:   publicKey,
	//	VerifyToken: verifyToken,
	//}
	//
	//fmt.Printf("> encrpytion request: %+v\n", encryptionReq)
	//
	//return c.reply(encryptionReq)

	loginRes := packet.LoginRes{
		Username: loginReq.Name,
		UUID:     loginReq.PlayerUUID,
	}

	fmt.Printf("> login response: %+v\n", loginRes)

	return c.reply(loginRes)
}

func (c *Client) handleEncryption(r *packet.Reader) error {
	var encryptionRes packet.EncryptionRes
	err := packet.Read(r, &encryptionRes)
	if err != nil {
		return err
	}

	fmt.Printf("< encryption response: %+v\n", encryptionRes)

	loginRes := packet.LoginRes{
		Username: c.player.Username,
		UUID:     c.player.UUID,
	}

	fmt.Printf("> login response: %+v\n", loginRes)

	return c.reply(loginRes)
}

func (c *Client) handleLoginAck() error {
	fmt.Println("< login acknowledged")

	c.state = packet.StateConfiguration

	return nil
}
