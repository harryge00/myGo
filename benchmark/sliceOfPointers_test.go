package main

import (
	"testing"
)

type MyStruct struct {
	F1, F2, F3, F4, F5, F6, F7 string
	I1, I2, I3, I4, I5, I6, I7 int64
}

func BenchmarkAppendingStructs(b *testing.B) {
	// var s []MyStruct
	s := make([]MyStruct, b.N)
	for i := 0; i < b.N; i++ {
		s = append(s, MyStruct{
			I1: int64(i),
		})
	}
}

func BenchmarkAppendingPointers(b *testing.B) {
	// var s []*MyStruct
	s := make([]*MyStruct, b.N)

	for i := 0; i < b.N; i++ {
		s = append(s, &MyStruct{
			I1: int64(i),
		})
	}
}
