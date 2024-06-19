package main

import (
	"fmt"
	"slices"
)

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	minPile, maxPile := 1, slices.Max(bloomDay)

	for minPile < maxPile {
		mid := (minPile + maxPile) / 2

		bouquets, size := 0, 0
		isPossible := false
		for _, j := range bloomDay {
			if j > mid {
				size = 0
				continue
			}
			size++
			if size >= k {
				bouquets++
				size = 0
				if bouquets == m {
					isPossible = true
					break
				}
			}
		}

		if isPossible {
			maxPile = mid
		} else {
			minPile = mid + 1
		}
	}

	return minPile
}

func main() {
	// 3
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	// -1
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 2))
	// 12
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
	// 1000000000
	fmt.Println(minDays([]int{1000000000}, 1, 1))
}
