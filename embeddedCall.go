package main

import (
	"fmt"
	"time"
)

type Mux struct {
	
}

func (m *Mux) Channel(source string) chan interface{} {

}

type AnotherMux struct {
	m mux
}

func (a *AnotherMux) Channel(source string) chan interface{} {
	c := a.m
	anotherCh := make(chan interface{})
	for {
		// do something to c
	}
	return anotherCh
}

func main() {
	fmt.Println("Go channels starting")

	ch := make(chan *myMsg) // unbuffered
	go chanSender(ch, "cs1")
	go chanSender(ch, "cs2")

	for msg := range ch {
		fmt.Println("Message", msg.seqNum, ":", msg.message)
	}
}

func chanSender(out chan<- *myMsg, prefix string) {
	seqNum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		out <- &myMsg{seqNum, fmt.Sprintf("%s: %s", prefix, "moo")}
		seqNum++
	}
	close(out)
}
