package main

import (
	"fmt"
	// "time"
)

func pumpEven(ch chan int, k int) {
	
	for i := 0; i < k; i += 2 {
		ch <- i
	}	
}

func pumpSingle(ch chan int, k int) {
	
	for i := 1; i < k; i += 2 {
		ch <- i
	}	
}


func main() {
	evenCh := make(chan int, 100)
	singleCh := make(chan int, 100)
	go pumpEven(evenCh, 200)
	go pumpSingle(singleCh, 200)
	LOOP:
	for {
		select {
		case a := <- evenCh:
			fmt.Println("Even", a, len(evenCh))
		case a := <- singleCh:
			fmt.Println("Single", a, len(singleCh))
		default:
			break LOOP
		}
	}

}
