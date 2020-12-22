package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	"strings"
)

func main() {
	abc := "aaaabb"
	fmt.Println(abc[0:3])
	arr1 := []string{"aaa", "bbb", "ccc"}
	str1 := strings.Join(arr1, ";")
	fmt.Println(str1)
	fmt.Println(strings.Split(str1, ";"))
	arr2 := []string{}
	fmt.Println("Shouldn't be nil", arr2 == nil)
	str2 := strings.Join(arr2, ";")
	fmt.Println(str2)
	fmt.Println(strings.Split(str2, ";"))

	fmt.Println(abc + arr1[2])

}
