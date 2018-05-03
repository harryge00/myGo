package main

import (
	"fmt"
	"strings"
)

func getStr() (string, int) {
	return "aaaa", 2
}
func multi() (a,b string){
	if 1 == 1 {
		a, err := getStr()
		fmt.Println(a, err)
		
	}
	return
}

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "d & u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println(multi())
}
