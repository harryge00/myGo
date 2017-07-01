package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

type a struct {
	b int
}

func main() {
	mmm := make(map[string]a)
	mmm["aa"] = a{b: 123}
	fmt.Println(mmm)
	d := &mmm["aa"]
	d.b = 456

	fmt.Println(mmm)

	m := map[string]float64{}
	var str string
	fmt.Println(str == "")
	str = "abc"
	fmt.Println("abc" == str)
	m["1"] = 1
	m["pi"] = 3.1415
	val, ok := m["ddd"]
	fmt.Printf("%v, value %g\n", ok, val)
	m2 := map[int64]string{}
	val2, ok2 := m2[567]
	fmt.Printf("%v, value %s\n", ok2, val2)
	var abc = "abc"
	var cde = "ab" + "c"
	fmt.Printf("%v,", abc == cde)
	var num1 int64
	num1 = 763422222222
	fmt.Printf("%d,\n", num1)
	times := make(map[int]int)
	times[2]++
	times[2]++
	fmt.Println(times[2])

	fmt.Println("")
	var gids map[int]bool
	gids = make(map[int]bool)
	gids[10] = true
	gids[21] = true
	for i := range gids {
		fmt.Println(i)
	}
	// for key, value := range m {
	// 	fmt.Printf("key %s, value %g\n", key, value)
	// }
	// m["exp"] = 2.7182
	// delete(m, "1")
	// for key, value := range m {
	// 	fmt.Printf("key %s, value %g\n", key, value)
	// }
	// m2 := map[string]bool{}
	// m2["abc"] =true
	// value, ok := m2["abc"]
	// fmt.Println("v,ok ",value,ok)
	// value2, ok2 := m2["abadfc"]
	// fmt.Println("v,ok ",value2,ok2)
}
