package main

import "fmt"

type tt struct {
	ss string
}

func main() {
	a("ABD")
}

var a = func(str string) {
	fmt.Printf("hl: %s\n", str)
	m := make(map[string]string)
	m["abc"] = "def"
	fmt.Printf("%s, %s\n", m["abc"], m["def"])
	if _, ok := m["def"]; !ok {
		fmt.Printf("not existed")
	}
	m2 := make(map[int]int)
	sss := tt{ss: "adffa"}
	fmt.Println(m2[11])
	fmt.Println(sss)
	m3 := make(map[int]*tt)
	m3[1] = &sss
	fmt.Println(m3[2])
}
