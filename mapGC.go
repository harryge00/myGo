package main

import (
    "log"
    "runtime"
    "time"
)

var intMap map[int]int
var cnt = 8192

func main() {
    printMemStats("start")

    initMap()

    printMemStats("map initialized")
    runtime.GC()
    printMemStats("First GC")


    log.Println(len(intMap))
    for i := 0; i < cnt; i++ {
        delete(intMap, i)
    }
    log.Println(len(intMap))

    runtime.GC()
    printMemStats("GC after deleting from map")

    intMap = nil
    time.Sleep(1 * time.Second)

    printMemStats("Map memory is not freed after assigning nil")

    runtime.GC()
    printMemStats("Finally")
}

func initMap() {
    intMap = make(map[int]int, cnt)

    for i := 0; i < cnt; i++ {
        intMap[i] = i
    }
}

func printMemStats(s string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    log.Printf("%s: Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", s, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
