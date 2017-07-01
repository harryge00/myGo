package main

import (
	// "encoding/json"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "test2.go")
	stdout, err := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	r1 := bufio.NewReader(stdout)
	bytes, isPrefix, err2 := r1.ReadLine()
	var buffer bytes.Buffer
	for err2 != io.EOF {
		if isPrefix {
			fmt.Print(string(bytes))
		} else {
			fmt.Println(string(bytes))
		}
		bytes, isPrefix, err2 = r1.ReadLine()
	}
	r2 := bufio.NewReader(stderr)
	bytes, isPrefix, err2 = r2.ReadLine()
	for err2 != io.EOF {
		if isPrefix {
			fmt.Print(string(bytes))
		} else {
			fmt.Println(string(bytes))
		}
		bytes, isPrefix, err2 = r2.ReadLine()
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
