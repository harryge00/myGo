package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

func pump(ch chan int) {
	for i := 0; i < 50; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
func main() {
	// 	ch1 := make(chan int)
	// 	go pump(ch1)
	// // pump hangs; we run
	// 	fmt.Println("channel res:")
	// 	for elem := range ch1 {

	// 		fmt.Println(elem)

	// 	}
	ch1 := make(chan int, 5)
	ch2 := ch1
	go pump(ch1)
	go suck(ch2)
	var input string
	fmt.Scanln(&input)
}
