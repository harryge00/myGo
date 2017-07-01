package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		i := 0
		for {
			c1 <- "from 1 " + strconv.Itoa(i)
			time.Sleep(time.Second * 2)
			i++
		}
	}()

	go func() {
		i := 0
		for {
			c2 <- "from 2 " + strconv.Itoa(i)
			time.Sleep(time.Second * 3)
			i++
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(20 * time.Second):
				fmt.Println(len(c1))
				fmt.Println(len(c2))
				break

			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
