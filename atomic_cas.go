package main

import "fmt"
import "time"
import "sync/atomic"

// HighWaterMark is a thread-safe object for tracking the maximum value seen
// for some quantity.
type HighWaterMark int64

// Update returns true if and only if 'current' is the highest value ever seen.
func (hwm *HighWaterMark) Update(current int64) bool {
    for {
        old := atomic.LoadInt64((*int64)(hwm))
        if current <= old {
            return false
        }
        if atomic.CompareAndSwapInt64((*int64)(hwm), old, current) {
            return true
        }
        fmt.Println(old, current, "interesting", atomic.LoadInt64((*int64)(hwm)))
    }
}

func tryUpdate(hwm *HighWaterMark, val int64) {
    // ok := hwm.Update(val)
    hwm.Update(val)
    // fmt.Println(val, ok)
}
func main() {

    var hwm HighWaterMark
    hwm = 1280
    fmt.Println(hwm)
    for i := 1; i < 10000; i++ {
        go tryUpdate(&hwm, int64(1280 + i))
    }
    time.Sleep(1 * time.Second)
    fmt.Println(hwm)
}

