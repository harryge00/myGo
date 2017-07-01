package main

import (
	"fmt"
	"os"
	"os/exec"
)

func print() {
	fmt.Println("output")
}

func main() {
	_, w, _ := os.Pipe()
	os.Stdout = w

	print()
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
