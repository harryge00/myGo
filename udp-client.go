package main
import (
    "fmt"
    "net"
    "bufio"
    "os"

)

func main() {
    p :=  make([]byte, 2048)
    fmt.Println(os.Args[1])
    conn, err := net.Dial("udp", os.Args[1] + ":9879")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}