package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("/home/tattoo/Codes/DisSys/mit2015/6.824/src/main/part_kjv12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	value := string(buf)
	result := strings.Fields(value)
	for _, w := range result {
		fmt.Print(w)
	}

}
