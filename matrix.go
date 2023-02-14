package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	// first there was a void and then
	// this matrix is created...
	// ⎡1  2  3⎤
	// ⎢4  5  6⎥
	// ⎣7  8  9⎦

	row1 := []float64{1, 2, 3}
	row2 := []float64{4, 5, 6}
	row3 := []float64{7, 8, 9}

	m := mat64.NewDense(3, 3, nil)
	m.SetRow(0, row1)
	m.SetRow(1, row2)
	m.SetRow(2, row3)

	// print all m elements
	fmt.Printf("m :\n%v\n\n", mat64.Formatted(m, mat64.Prefix(""), mat64.Excerpt(0)))

	// then followed by the matrix transpose
	mT := m.T()

	// print all mT elements
	fmt.Printf("mT :\n%v\n\n", mat64.Formatted(mT, mat64.Prefix(""), mat64.Excerpt(0)))

	// the matrices decided to go forth and multiply...
	//mX := m * mT

	mX := mat64.NewDense(3, 3, nil) // a new nil matrix of 3 x 3
	mX.MulElem(m, mT)

	// print all mX elements
	fmt.Printf("mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))

	// and add another matrix to make another family member...
	mA := mat64.NewDense(3, 3, nil)
	mA.Add(mX, mT)

	// print all mA elements
	fmt.Printf("mA :\n%v\n", mat64.Formatted(mA, mat64.Prefix(""), mat64.Excerpt(0)))
	fmt.Println("    =    ")
	fmt.Printf("mX :\n%v\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
	fmt.Println("    +    ")
	fmt.Printf("mT :\n%v\n\n", mat64.Formatted(mT, mat64.Prefix(""), mat64.Excerpt(0)))

	// in order to be more fruitful, first...one must find the determination (determinant)

	determinant := mat64.Det(mX)
	fmt.Println("Determinant of mX : ", determinant)

	// with strong determination (determinant), then linear algebra operations can be solved ....

	err := mX.Solve(mX, m)
	fmt.Println("Any error? : ", err)

	fmt.Printf("Solved mX :\n%v\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
}
