package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Grade int
	Level int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return (s[i].Level > s[j].Level) || (s[i].Grade < s[j].Grade)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fmt.Println("Sorting Demo in Go")

	students := []Student{
		{"Denis", 12, 0},
		{"Ascolin", 50, 33},
		{"IIsus", 99, 0},
		{"Mario", 2, 0},
		{"Gogaie", 37, 22},
		{"Dentistul", 76, 0},
	}

	fmt.Println(students)
	fmt.Println(sort.IsSorted(Students(students)))

	sort.Sort(Students(students))

	fmt.Println(students)
	// fmt.Println(sort.IsSorted(students))

	// sort.Sort(sort.Reverse(students))
	// fmt.Println(students)
}
