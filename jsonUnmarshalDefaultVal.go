package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	A int
	B string
	C string
}

var example []byte = []byte(`{"A": 1, "C": "3"}`)

func main() {
	out := Test{
		A: 22,
		B: "default b",
		// C will be ""
	}
	out2 := out
	err := json.Unmarshal(example, &out2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v", out, out2)
}
