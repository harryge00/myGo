package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

func pump(ch chan int) {
	fmt.Println(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 90; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println("suck", v, len(ch))
		}
	}()
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
	suck(ch1)
	for i := 0; i < 6; i++ {
		// fmt.Println(i)
		ch1 <- i
	}
	stream := pump2()
	// for {
	// 	value, ok:= <- stream
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(value)
	// }
	suck(stream)
	var x int
	fmt.Scanln(&x)
	var dd *int
	dd = &x
	fmt.Println(dd)
}
