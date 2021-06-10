package main

import "fmt"
import "strconv"

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func main() {
    toleration := 10
    f1 := (1 - float64(toleration)/100)
    fmt.Println("f1 should be 0.9: ", f1)
    fmt.Println(FloatToString(21312421.213123))
}
