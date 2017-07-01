package main

import (
	"fmt"
)

type stu struct {
	id   int
	name string
}

func main() {
	classname := "15-440" // Type inference

	fmt.Println("Hello,", classname)
	a, b, c := 0, 0, 0
	fmt.Println("%d %d %d", a, b, c)
	greeting := "你好"
	var str string
	if str == "" {
		fmt.Println(greeting, classname)
	}
	stu1 := &stu{1, "wang han"}
	fmt.Println(stu1.name)
}
