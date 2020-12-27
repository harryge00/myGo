package main

import (
	"fmt"
)

func main() {
	var counter int = 1

	func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function")
	}("Ricky")

	funcVar := func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function assigned to a variable.")
	}

	funcVar("Johnny")

	closure := func(str string) {
		fmt.Println("Hi", str, "I'm a closure.")
		for i := 1; i < 5; i++ {
			fmt.Println("Counter incremented to:", counter)
			counter++
		}
	}

	fmt.Println("Counter is", counter, "before calling closure.")
	closure("Sandy")
	fmt.Println("Counter is", counter, "after calling closure.")

	var recursion func(j int)
	recursion = func(i int) {
		if i >= 10 {
			fmt.Println("Break Traversal now")
			return
		}
		fmt.Printf("Current i: %v\n", i)
		recursion(i + 1)
	}
	recursion(counter)
}
