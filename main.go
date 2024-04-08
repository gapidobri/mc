package main

import (
	"fmt"
	"log"
	cl "mc/client"
	"net"
)

func main() {

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
		client := cl.NewClient(conn)
		go client.HandleRequest()
	}
}
