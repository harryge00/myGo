package main

import "fmt"

type Config struct {
	Num    int                // config number
	Shards [10]int64          // gid
	Groups map[int64][]string // gid -> servers[]
}

func main() {
	var config Config
	config.Num = 100
	config.Groups = make(map[int64][]string)
	config.Shards[1] = 2
	fmt.Println(config.Shards)
	arr := make([]int, 0, 100)
	for i := 0; i < 5; i++ {
		arr[len(arr)] = i + 100
	}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
