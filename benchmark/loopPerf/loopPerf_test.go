package main

import (
	"testing"
)

func createMatrix(x, y int) [][]int {
	m := make([][]int, x)
	for i := 0; i < x; i++ {
		m[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			m[i][j] = i*y + j
		}
	}
	return m
}

func add1(m [][]int, x, y int) int {
	s := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			s += m[i][j]
		}
	}
	return s
}

func add2(m [][]int, x, y int) int {
	s := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			s += m[j][i]
		}
	}
	return s
}

func BenchmarkAdd1(b *testing.B) {
	m := createMatrix(1000, 1000)
	for i := 0; i < b.N; i++ {
		add1(m, 1000, 1000)
	}
}

func BenchmarkAdd2(b *testing.B) {
	m := createMatrix(1000, 1000)
	for i := 0; i < b.N; i++ {
		add2(m, 1000, 1000)
	}
}

