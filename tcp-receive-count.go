package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"fmt"
)

func main() {
	port := flag.Int("port", 15333, "Port to accept connections on.")
	//host := flag.String("host", "127.0.0.1", "Host or IP to bind to")
	flag.Parse()

	l, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections on port", strconv.Itoa(*port))
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")
	totalSize := 0
	for {
		buf := make([]byte, 4096)
		size, err := conn.Read(buf)
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Received %d bytes", totalSize)))
			return
		}
		// data := buf[:size]
		// log.Println("Read new data from connection", data)
		// conn.Write(data)
		totalSize += size
	}
}
