package main

import "fmt"

type Shape interface {
  area() float64
}

type Square struct {
  x1, y1, x2, y2 float64
}

func (s *Square) area() float64 {
  l := distance(s.x1, s.y1, s.x2, s.y2) // calculate distance between 2 points
  return l * l
}

func distance(a, b, c, d float64) float64 {
	return 2.01
}

func main() {
	s := &Square{1, 2,3,4}
	fmt.Println(s.area())
}