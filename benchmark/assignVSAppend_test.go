package main

import "testing"

var result []int

const size = 1000

const iterations = 1000 * 1000

func doAssign() {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	result = data
}

func doAppend() {
	var data []int
	for i := 0; i < size; i++ {
		data = append(data, i)
	}
	result = data
}

func BenchmarkAssign(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		doAssign()
	}
}

func BenchmarkAppend(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		doAppend()
	}
}
