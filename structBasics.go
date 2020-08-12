package main

import "fmt"

type T struct{
	name string
}

type Dog struct {
	age  int
	name string
}

func main() {
	aaa := T{name:"aaa"}
	bbb := aaa
	fmt.Println("aaa==bbb?", aaa == bbb)
	
	bbb.name="bbb"

	fmt.Println(aaa, bbb)
	fmt.Println(&aaa, &bbb)
	arr := []T{aaa, bbb}
	for _, t := range arr {
		t.name = "xxx"
	}
	fmt.Println(arr)
	fmt.Println(aaa, bbb)
	fmt.Println("aaa==bbb?", aaa == bbb)

	roger := Dog{5, "Roger"}
	mydog := roger
	if roger == mydog {
		fmt.Println("Roger and mydog are equal structs")
	}
}
