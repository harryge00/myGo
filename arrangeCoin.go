package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	// (1 + k) * k / 2 <= n
	// (1+k + 1) * (k+1)/2 > n
	// (1+k)*k <= 2n
	// (2 + k) * (k + 1) > 2n
	// 1 2 3 4 4
	// 1 2 3 4 3
	n2 := 2 * n
	root := math.Sqrt(float64(n2))
	r := int(root)
	// fmt.Println(root, r)
	if r*(r+1) > n2 {
		return r - 1
	}
	return r
}

func main() {
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(13))
	fmt.Println(arrangeCoins(14))
	fmt.Println(arrangeCoins(15))
}
