package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2022-12-12 18:22:09"
	date, error := time.Parse("2006-01-02 15:04:05", dateString)

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("Type of dateString: %T\n", dateString)
	fmt.Printf("Type of date: %T\n", date)
	fmt.Println()
	fmt.Printf("Value of dateString: %v\n", dateString)
	fmt.Printf("Value of date: %v\n", date)
	fmt.Printf("Value of date: %v", date.Unix())
}
