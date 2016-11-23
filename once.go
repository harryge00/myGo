package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("another")
			})
			once.Do(func() {
				fmt.Println(i)
			})
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
