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
	if a,b = 3, 4; str == "" {
		fmt.Println(greeting, classname)
		fmt.Println(a, b)
	}
	stu1 := &stu{1, "wang han"}
	fmt.Println(stu1.name)
}
