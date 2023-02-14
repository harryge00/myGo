package main

import "testing"

type PAdder struct {
	x, y int
}

func (a *PAdder) add() int {
	return a.x + a.y
}

type VAdder struct {
	x, y int
}

func (a VAdder) add() int {
	return a.x + a.y
}

func BenchmarkPAdder(b *testing.B) {
	accum := 0
	adder := PAdder{
		x: 100,
		y: 999,
	}
	for i := 0; i < b.N; i++ {
		accum = adder.add()
	}
	_ = accum
}

func BenchmarkVAdder(b *testing.B) {
	accum := 0
	adder := VAdder{
		x: 100,
		y: 999,
	}
	for i := 0; i < b.N; i++ {
		accum = adder.add()
	}
	_ = accum
}
