package main

import (
	"fmt"
	"strconv"
	"time"
)

type PaxosArgs struct {
	Seq int
	Num int
	Val interface{}
}

func change(arg PaxosArgs) {
	arg.Seq = 666
	arg.Val = "youyouyou"
	fmt.Println(arg.Seq, ":", arg.Val)
}
func main() {
	var v interface{}
	fmt.Println(v == nil)
	// v="you are cool"
	arg := PaxosArgs{5, 28, "you are handsome"}
	fmt.Println(arg.Val)
	change(arg)
	fmt.Println(arg.Seq, ":", arg.Val)
	for i := 0; i < 15; i++ {
		str := strconv.FormatInt(time.Now().UnixNano(), 10)
		fmt.Println(str)
	}
	dd := "d+d"
	dd2 := "d"
	dd2 = dd2 + "+d"
	fmt.Printf("%s==%s?%v", dd, dd2, dd == dd2)
}
