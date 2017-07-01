package main

import (
	"fmt"
	//"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("ping snet", i)
		c <- "ping"
	}
}

/*func ponger(c chan string) {
  for i := 0; ; i++ {
  fmt.Println("pong sent",i)
    c <- "pong"
  }
}*/
func printer(c chan string) {
	for {
		fmt.Println(" print ", <-c)
		//time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	//go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
