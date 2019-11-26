package main

import (
	"fmt"
	"os"
)

func main() {
	dd := os.Getenv("abcd")
	fmt.Println(dd)
}
