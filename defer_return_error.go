// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func check(str string) (err error) {
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
	log.Println("executing")
	return testOpen(str)
}

func testOpen(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	s := string(buf.Bytes())
	fmt.Printf("%s", s)
	return nil
}

func main() {
	err := check("sbsb")
	fmt.Printf("check res:%v\n", err)
}
