package main

import (
	"fmt"
	"time"
)

func pump(ch chan int, k int) {
	
	for i := 0; i < k; i++ {
		ch <- i
	}	
}


func main() {
	ch := make(chan int, 100)
	go pump(ch, 200)
	time.Sleep(1 * time.Millisecond)
	LOOP:
	for {
		select {
		case a := <- ch:
			fmt.Println(a, len(ch))
		default:
			fmt.Println("quit", len(ch))
			break LOOP
		}
	}

}
