package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	sum := 0
	for i := 0; i < len(seats); i++ {
		t := seats[i] - students[i]
		if t < 0 {
			t *= -1
		}

		sum += t
	}

	return sum
}

func main() {
	// 4
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	fmt.Println(minMovesToSeat(seats, students))
}
