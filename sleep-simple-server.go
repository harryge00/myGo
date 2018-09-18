package main

import (
	"log"
	"io"
	"net"
	"time"
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.Header)
	log.Println(r.RemoteAddr)
	// RemoteAddr

	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	log.Println("Hello, 你好!")
	time.Sleep(3 * time.Second )
	io.WriteString(w, "Hello, 你好!")
}

func udpServer() {
	//Basic variables
	port := ":9879"
	protocol := "udp"

	//Build the address
	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		panic(err)
	}


	//Create the connection
	udpConn, err := net.ListenUDP("udp",  udpAddr)
	if err != nil {
		panic(err)
	}

	//Keep calling this function
	for {
		display(udpConn)
	}

}
func display(conn *net.UDPConn) {

	var buf [2048]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error Reading")
		return
	} else {
		fmt.Println(string(buf[0:n]))
		fmt.Println("Package Done")
	}

}

func main() {
	go udpServer()
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9878", nil)
}
