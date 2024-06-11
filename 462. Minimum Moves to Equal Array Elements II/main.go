package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}

	sort.Ints(nums)
	mid := nums[len(nums)/2]

	count := 0
	for _, num := range nums {
		if num > mid {
			count += num - mid
		} else {
			count += mid - num
		}
	}

	return count
}

func main() {
	// 2
	nums := []int{1, 2, 3}
	fmt.Println(minMoves2(nums))
}
