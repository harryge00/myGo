
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_CONN_NUM = 5
)

func main() {
	udpListener, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 2222,
	})
	checkError(err)
	tcpListener, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 2223,
	})
	checkError(err)
	defer udpListener.Close()
	defer tcpListener.Close()

	//
	// tcp
	//
	var curTcpConnNum int = 0
	connTcpNumChan := make(chan net.Conn)
	connTcpNumAdjustChan := make(chan int)

	// 调整现在的链接数
	go func() {
		for numAdjust := range connTcpNumAdjustChan {
			curTcpConnNum += numAdjust
		}
	}()

	go func() {
		for _ = range time.Tick(1e9) {
			fmt.Printf("当前链接数: %d\n", curTcpConnNum)
		}
	}()

	// TODO: 只有在用户请求时才创建协程
	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range connTcpNumChan {
				connTcpNumAdjustChan <- 1
				EchoFunc(conn)
				connTcpNumAdjustChan <- -1
			}
		}()
	}

	//
	// tcp && udp
	//
	for {
		print("执行到了1\n")

		print("执行到了3\n")
		conn, err := tcpListener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		connTcpNumChan <- conn // 有链接进来了

		print("执行到了4\n")
	}

	for {
		handleClient(udpListener)
	}
}

func handleClient(conn *net.UDPConn) {
	print("执行到了2\n")
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	print("Client:%s", string(buf[:10]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func EchoFunc(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		checkError(err)

		_, err = conn.Write(buf)
		checkError(err)
	}
}
