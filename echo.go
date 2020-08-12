
package main

import (
    "bufio"
    "log"
    "net"
    "encoding/hex"
    "io/ioutil"
)

func main() {
    log.Printf("echo-server listening on tcp port 8800") 
    
    ln, err := net.Listen("tcp", ":8800") 
    if err != nil { 
        log.Fatalf("listen error, err=%s", err) 
    } 
    
    accepted := 0 
    for { 
            conn, err := ln.Accept() 
            if err != nil { 
                log.Fatalf("accept error, err=%s", err) 
            } 
            accepted++ 
            go handleConnection2(conn)
            log.Printf("connection accepted %d", accepted) 
    } 
} 

func handleConnection2(conn net.Conn) {
    bufr := bufio.NewReader(conn)
    body, _ := ioutil.ReadAll(bufr)
    log.Println(string(body))
    conn.Write([]byte("ok"))
    conn.Close()
}

func handleConnection(conn net.Conn) { 
    bufr := bufio.NewReader(conn)
    buf := make([]byte, 1024)
    
    for {
        readBytes, err := bufr.Read(buf)
        if err != nil {
            log.Printf("handle connection error, err=%s", err)
            conn.Close()
            return
        }
        log.Printf("<->\n%s", hex.Dump(buf[:readBytes]))
        conn.Write([]byte("CLOUDWALK "+string(buf[:readBytes])))
    }

}