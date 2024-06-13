package main

import "fmt"

func longestSubarray(nums []int) int {
	maxLen := 0
	count := 0
	found := false
	newStart := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count += 1
			continue
		}

		if found {
			if count > maxLen {
				maxLen = count
			}
		} else {
			found = true
			newStart = i
			continue
		}

		count = 0
		found = false
		i = newStart
	}

	if found {
		if count > maxLen {
			maxLen = count
		}

		return maxLen
	}

	return count - 1
}

func main() {
	// 4
	nums := []int{0, 1, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))

	// 5
	nums = []int{1, 1, 1, 1, 1, 1}
	fmt.Println(longestSubarray(nums))

	// 1
	nums = []int{0, 0, 0, 0, 0, 1}
	fmt.Println(longestSubarray(nums))

	// 5
	nums = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
