package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	//	"os"
	"path/filepath"
)

func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	ln := ""
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)
}
