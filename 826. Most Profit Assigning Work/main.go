package main

import (
	"fmt"
	"sort"
)

type DiffProfit struct {
	Difficulty int
	Profit     int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	diffProfit := make([]DiffProfit, 0, len(profit))
	for i := range profit {
		diffProfit = append(diffProfit, DiffProfit{
			Difficulty: difficulty[i],
			Profit:     profit[i],
		})
	}

	sort.Slice(diffProfit, func(i, j int) bool {
		if diffProfit[i].Difficulty == diffProfit[j].Difficulty {
			return diffProfit[i].Profit < diffProfit[j].Profit
		}

		return diffProfit[i].Difficulty < diffProfit[j].Difficulty
	})
	sort.Ints(worker)

	bestProfit := 0
	diffIndex := 0
	s := 0
	for _, w := range worker {
		j := diffIndex
		for j < len(diffProfit) && w >= diffProfit[j].Difficulty {
			if bestProfit < diffProfit[j].Profit {
				bestProfit = diffProfit[j].Profit
			}
			j += 1
		}

		s += bestProfit
	}

	return s
}

func main() {
	// 100
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	fmt.Println(maxProfitAssignment(difficulty, profit, worker))
}
