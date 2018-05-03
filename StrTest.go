package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var str string
	str += "abc"
	fmt.Println(str)
	arr2d := make([][]string, 3)
	arr := []string{"a", "b"}
	arr2d[0] = arr
	arr[1] = "c"
	var wg sync.WaitGroup
	arr2d[1] = arr
	fmt.Println(arr2d)
	name := "zk-cluster-2"
	i := strings.LastIndex(name, "-")
	wg.Wait()
	fmt.Println(name[i+1:])
}
