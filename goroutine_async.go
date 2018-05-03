package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	arr := []int{13, 48, 4,5,6, 7,8,9, 10}
	for i, num := range arr {
		go func(i int, num *int) {
			fmt.Println(i, *num)
		} (i, &num)
	}
	var input string
	fmt.Scanln(&input)
}
