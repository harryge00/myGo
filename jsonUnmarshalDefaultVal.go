package main

import (
	"encoding/json"
	"fmt"
)

type Metrics struct {
	Tags  string  `json:"Tags,omitempty"`
	Tests []*Test `json:"Tests,omitempty"`
}
type Test struct {
	A int    `json:"A,omitempty"`
	B string `json:"B,omitempty"`
	C string `json:"C,omitempty"`
}

var example []byte = []byte(`{"A": 1, "C": "3"}`)
var example2 []byte = []byte(`{}`)
var jsonArr []byte = []byte(`{"Tags":"abcd","Tests":[{"A": 1, "C": "3"}]}`)
var jsonArr2 []byte = []byte(`{"Tags":"abcd","Tests":[{"B": "bb"}]}`)
var emptyArr []byte = []byte(`{"Tags":"abcd","Tests":[]}`)

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
	fmt.Printf("%+v %+v\n", out, out2)

	err = json.Unmarshal(example2, &out2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", out, out2)
	// originalArr := []Test{}
	oldMetrics := Metrics{}
	fmt.Println("arr", json.Unmarshal(jsonArr, &oldMetrics))
	fmt.Println(oldMetrics)
	fmt.Println(oldMetrics.Tests[0])
	fmt.Println(json.Unmarshal(jsonArr2, &oldMetrics))
	fmt.Println(oldMetrics)
	fmt.Println(json.Unmarshal(emptyArr, &oldMetrics))
	fmt.Println(oldMetrics)
}
