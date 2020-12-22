package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func calculateSum(s []int, c chan []int) {
	for i := 1; i < len(s); i++ {
		s[i] += s[i - 1]
	}
	c <- s
}
func main() {
	s := []int{1,2,3,4,5,6}
	c := make(chan []int)
	go calculateSum(s[:len(s)/2], c)
	go calculateSum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(s)
	fmt.Println(x, y)
	close(c)
	testVal, ok := <- c
	fmt.Println(testVal, ok)
	var a []int
	var m map[int]int
	fmt.Println(a, m)

	testCloseIntChan()
}

func testCloseIntChan() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	fmt.Println("shoulde be 1", <-c) //1
	fmt.Println("shoulde be 2", <-c) //1
	fmt.Println("shoulde be 0", <-c) //1
	fmt.Println("shoulde be 0", <-c) //1
}