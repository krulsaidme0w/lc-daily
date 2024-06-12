package main

import "fmt"

func sortColors(nums []int) {
	colorsCount := make(map[int]int, 0)

	for _, color := range nums {
		if _, ok := colorsCount[color]; ok {
			colorsCount[color] = colorsCount[color] + 1
		} else {
			colorsCount[color] = 1
		}
	}

	start := 0

	for i := 0; i < colorsCount[0]; i++ {
		nums[start] = 0
		start += 1
	}

	for i := 0; i < colorsCount[1]; i++ {
		nums[start] = 1
		start += 1
	}

	for i := 0; i < colorsCount[2]; i++ {
		nums[start] = 2
		start += 1
	}
}

func main() {
	// [0,0,1,1,2,2]
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
