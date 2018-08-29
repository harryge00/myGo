package main
import (
   "log"
    "net"
    "os"

    "golang.org/x/net/icmp"
    "golang.org/x/net/ipv4"
)

func main() {
    wm := icmp.Message{
        Type: ipv4.ICMPTypeEcho, Code: 0,
        Body: &icmp.Echo{
            ID: 1, Seq: 1,
            Data: []byte("HELLO-R-U-THERE"),
        },
    }
    wb, err := wm.Marshal(nil)

   	if err != nil {
        log.Fatal(err)
    }

    conn, err := net.Dial("ip4:icmp", "10.30.96.99") 
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write(wb)
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("write to server = ", string(wb))

    reply := make([]byte, 512)

    _, err = conn.Read(reply)
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("reply from server=", string(reply))

    conn.Close()
}