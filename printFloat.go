package main

import "fmt"
import "strconv"

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func main() {
    fmt.Println(FloatToString(21312421.213123))
}
