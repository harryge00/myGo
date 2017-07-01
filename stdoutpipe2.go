package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	// "bufio"
	"bytes"
)

func main() {
	cmd := exec.Command("bash", "test2.sh")
	var buf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("all:")
	fmt.Print(buf.String())
}
