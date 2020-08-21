package main

import "fmt"
import "runtime"
import "time"
import "sync/atomic"

func main() {
        fmt.Println(runtime.GOMAXPROCS(1))
        fmt.Println(runtime.GOMAXPROCS(1))
        // fmt.Println(runtime.NumCPU())
        var ops uint64 = 0 
        go func() {
                go func() {
                        fmt.Println("inner loop")
                        for {
                                atomic.AddUint64(&ops, 1)
                        }   
                } ()
                fmt.Println("outer loop")
                opsFinal := atomic.LoadUint64(&ops)
                fmt.Println("outer ops:", opsFinal)
                
        }()
        time.Sleep(time.Second)
        opsFinal := atomic.LoadUint64(&ops)
        fmt.Println("ops:", opsFinal)
}
