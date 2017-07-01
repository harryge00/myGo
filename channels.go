package main

import "log"
import "fmt"
import "time"

func sum(a []int, c chan int) {
	// log.Println(len(a))
	sum := 0
	for _, v := range a {
		// log.Println(v,sum)
		sum += v
	}
	log.Println("sum ", sum)
	c <- sum // send sum to c
	log.Println("len: ", len(c))
}

func main() {
	a := []int{5, 4, 3, 2, 1, 0}
	var c = make(chan int)
	for i := 1; i < 6; i++ {
		go sum(a[:i], c)
	}
	// go sum(a[len(a)/2:], c)
	log.Println("len: ", len(c))
	go func() {
		time.Sleep(500)
		i := 1
		for {
			x := <-c
			fmt.Println(x, i)
			i++
		}
	}()

	var input string
	fmt.Scanln(&input)
}
