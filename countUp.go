// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import "fmt"

var count int

func addCount(n int) {
	for i := 0; i < n; i++ {
		go func() {
			count++
		}()
	}
}
func main() {
	count = 0
	for j := 0; j < 10; j++ {
		go addCount(4000)
	}
	var text string
	fmt.Scanln(&text)
	fmt.Println(count)
}
