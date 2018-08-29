package main

import "fmt"
import "time"

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	for i := range c {
		fmt.Println(i)
	}
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		// close(c)
	}()
	fmt.Println("Finished")
}
