// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("sbcdf".Substring(1, 3))
	fmt.Printf("%v\n", time.Hour*24)
	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Open("dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	s := string(buf.Bytes())
	fmt.Printf("%s", s)

}
