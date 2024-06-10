package main

import "fmt"

func maxDistToClosest(seats []int) int {
	max := -1

	left := -1
	count := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			count++
			continue
		}

		if left == -1 {
			if max < count {
				max = count
			}
		} else if max < (count+1)/2 {
			max = (count + 1) / 2
		}

		left = i
		count = 0
	}

	if max < count {
		max = len(seats) - left - 1
	}

	return max
}

func main() {
	// 2
	seats := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(seats))

	// 4
	seats = []int{0, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(seats))

	// 3
	seats = []int{1, 0, 0, 0}
	fmt.Println(maxDistToClosest(seats))
}
