package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	result := make([]string, 0, 0)
	start := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		prev := nums[i-1]

		if current-prev != 1 {
			if count == 1 {
				result = append(result, strconv.Itoa(start))
			} else {
				t1 := strconv.Itoa(start)
				t2 := strconv.Itoa(prev)
				result = append(result, t1+"->"+t2)
			}
			start = nums[i]
			count = 1
		} else {
			count += 1
		}
	}

	if count == 1 {
		result = append(result, strconv.Itoa(nums[len(nums)-1]))
	} else {
		t1 := strconv.Itoa(start)
		t2 := strconv.Itoa(nums[len(nums)-1])
		result = append(result, t1+"->"+t2)
	}

	return result
}

func main() {
	// ["0->2","4->5","7"]
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
