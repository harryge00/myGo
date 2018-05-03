package main

import (
	"fmt"
	"time"
	"k8s.io/apimachinery/pkg/util/wait"
)

type myMsg struct {
	seqNum  int
	message string
}



func main() {
	fmt.Println("Go channels starting")	
	ch := make(chan struct{})
	i := 1
	go func() {
		wait.Until(func() {
			fmt.Println("do", i)
			i++
			}, 500 * time.Millisecond, ch)
	} ()
	time.Sleep(3 * time.Second)
	var s struct{}
	ch <- s
}
