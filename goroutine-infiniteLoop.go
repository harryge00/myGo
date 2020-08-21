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
        for i := 0; i < 50; i++ {
                go func() {
                        for {
                                atomic.AddUint64(&ops, 1)
                                // fmt.Print()
                        }
                }()
        }
        time.Sleep(time.Second)
        opsFinal := atomic.LoadUint64(&ops)
        fmt.Println("ops:", opsFinal)
}
