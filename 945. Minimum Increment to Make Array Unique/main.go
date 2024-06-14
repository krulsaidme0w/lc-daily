package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	minIncrement := 0

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		prev := nums[i-1]

		if curr <= prev {
			increment := prev - curr + 1
			minIncrement += increment

			nums[i] = prev + 1
		}
	}

	return minIncrement
}

func main() {
	// 1
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))

	// 6
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
