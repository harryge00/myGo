package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := []int{9, 8, 7, 6}
	b := a
	fmt.Println(a, b)
	b[1] = 999
	fmt.Println(a, b)
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&b[0]), unsafe.Pointer(&b[1]), unsafe.Pointer(&b[2]))
	fmt.Println(unsafe.Pointer(&a[0]), unsafe.Pointer(&a[1]), unsafe.Pointer(&a[2]))

	c := a[1:]
	fmt.Println(unsafe.Pointer(&c[0]), unsafe.Pointer(&c[1]), unsafe.Pointer(&c[2]))


	b = append(b, 111)
	fmt.Println(a, b)

	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Pointer(&b[0]), unsafe.Pointer(&b[1]), unsafe.Pointer(&b[2]))
	fmt.Println(unsafe.Pointer(&a[0]))

	b[0] = 0
	fmt.Println(a, b)



}
