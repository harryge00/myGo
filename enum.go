// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"fmt"
)

type ScalingState int

const (
	SameScale ScalingState = iota
	ScaleUp
	ScaleDown
	Mixed
)

func (s ScalingState) String() string {
	return [...]string{"SameScale", "ScaleUp", "ScaleDown", "Mixed"}[s]
}

func main() {
	s1 := Mixed
	fmt.Println(s1 - 1)
}
