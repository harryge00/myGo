package main

import (
	"os"
	// "bufio"
	"fmt"
	// "io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	bb := make([]byte, 100)
	cmd.Run()
	os.Stdout.Read(bb)
	// r := bufio.NewReader()
	// line, _, _ := r.ReadLine()
	fmt.Printf("s:%s", string(bb))
}
