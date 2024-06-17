package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		s := left*left + right*right
		if s == c {
			return true
		}

		if s > c {
			right -= 1
		} else {
			left += 1
		}
	}

	return false
}

func main() {
	// true
	fmt.Println(judgeSquareSum(5))

	// true
	fmt.Println(judgeSquareSum(100000))
}
