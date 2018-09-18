package main

import (
	"fmt"
	"time"
)

func sleep() {
	fmt.Println("sleeping")
	time.Sleep(time.Second * 1)
	fmt.Println("Waked up")
}

func f() {
	go sleep()
}
func main() {
	go f()
	var input string
	fmt.Scanln(&input)
}
