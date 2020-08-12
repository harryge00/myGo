package main

import "fmt"

func arrayModify(array [5]int) {
    fmt.Printf("array pointer in fun: %p\n", &array)
    newArray := array
    newArray[0] = 88
}
func sliceModify(slice []int) {
    fmt.Printf("slice pointer in fun: %p\n", slice)
    newSlice := slice
    newSlice[0] = 88
}
func main() {
    array := [5]int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(len(array), cap(array))
    fmt.Println(len(slice), cap(slice))
    subSlice := slice[1:3]
    subArray := array[1:3]
    fmt.Println("subSlice ", subSlice)
    fmt.Println("subArray ", subArray)
    fmt.Printf("subArray %p &array[1] %p\n", subArray, &array[1])
    fmt.Printf("array pointer in main: %p\n", &array)
    fmt.Printf("slice pointer in main: %p\n", slice)

    arrayModify(array)
    sliceModify(slice)
    fmt.Println("After modifying array", array)
    fmt.Println("After modifying slice", slice)

    fmt.Println("subSlice ", subSlice, len(subSlice), cap(subSlice))
    fmt.Println("subArray ", subArray, len(subArray), cap(subArray))
    sl2 := append(slice, 6)
    fmt.Println(len(sl2), cap(sl2), sl2, slice)
    slice[0] = 999
    fmt.Println("Append makes a new slice", slice, sl2)
    fmt.Printf("Append makes a new slice %p %p\n", slice, sl2)

    var foo []int
    foo = make([]int, 5, 8)
    foo[3] = 42
    foo[4] = 100
    fmt.Println("foo", foo)
    
    foo[4] = foo[4] >> 4
    fmt.Println("foo", foo)
    foo = append(foo, make([]int, 99)...)
    foo[90] = 9999
    fmt.Println("Resized array: ", foo)


    bar  := foo[1:4]
    bar[1] = 99
    fmt.Println(foo,bar)
    fmt.Println(len(foo),cap(foo))
    fmt.Println(len(bar),cap(bar))
    foo2 := append(foo, 33)
    foo2[1] = 77
    fmt.Println(foo, len(foo), cap(foo))
    fmt.Println(foo2, len(foo2), cap(foo2))

}