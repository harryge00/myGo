package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	if err != nil {
	    fmt.Println(err)
	}
	fmt.Println(t.String())
	t2 := t.Add(5 * time.Second)
	fmt.Println(t.Format("2006-01-02T15:04:05.000Z"), t2.Format("2006-01-02T15:04:05.000Z"))
}
