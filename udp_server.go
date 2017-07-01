package main

import (
	"log"
	"net"
	"strconv"
)

const (
	BUFFSIZE = 1024
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		log.Println("[server] resolving error ", err)
		return
	}
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Println("[server] listening error ", err)
		return
	}
	readFromClient(conn)
}
func readFromClient(conn *net.UDPConn) {
	var buffer [BUFFSIZE]byte
	i := 0
	for {
		_, addr, err := conn.ReadFromUDP(buffer[0:])
		if err != nil {
			log.Println("[server] ReadFromUDP error ", err)
			continue
		}
		log.Println(i, string(buffer[0:]))
		str := "No " + strconv.Itoa(i) + " msg received."
		conn.WriteToUDP([]byte(str), addr)
		i++
	}
}
