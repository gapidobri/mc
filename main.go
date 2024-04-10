package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	cl "mc/client"
	"net"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":25565")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening for connections")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := cl.NewClient(conn, key)
		go client.HandleRequest()
	}
}
