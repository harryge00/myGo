package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		log.Println("[server] resolving error ", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("DialUDP Error", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for { // read in input from stdin
		log.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		log.Println("Sending ", text) // listen for reply
		if _, err := conn.Write([]byte(text)); err != nil {
			log.Print("Write error ")
			return
		}
		var buffer [1024]byte
		if _, err := conn.Read(buffer[0:]); err != nil {
			log.Print("Read error ")
			return
		}
		log.Println("Message from server: ", string(buffer[0:]))
	}
}
