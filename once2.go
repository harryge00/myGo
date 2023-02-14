package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func doOnce(i int) {
	once.Do(func() {
		fmt.Println("once", i)
	})

}
func main() {
	doOnce(0)
	doOnce(1)
	doOnce(2)
	fmt.Println("all done")
}
