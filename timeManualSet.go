package main

import (
	"fmt"
	"time"
)

func main() {
	//time.Now() example
	t := time.Now()
	fmt.Println(t)

	//time.Date() Example
	t = time.Date(t.Year(), t.Month(), t.Day(), 23, 0, 0, t.Nanosecond(), t.Location())
	fmt.Println(t)

	//time.Parse() Example
	//Parse YYYY-MM-DD
	t, _ = time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(t)
}
