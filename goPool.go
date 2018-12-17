package main

import (
	"fmt"
	//"time"
)

func pushCh(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		passParam := i
		go func(param int) {
			fmt.Println(passParam, param)
		}(passParam)
	}
	go pushCh(c)
	go func() {
		for xx := range c {
			fmt.Println(xx)
		}
		fmt.Println("Done")
	} ()


	var input string
	fmt.Scanln(&input)
	close(c)
	fmt.Scanln(&input)
	
}
