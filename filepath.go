package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.EvalSymlinks(os.Args[1]))
}
