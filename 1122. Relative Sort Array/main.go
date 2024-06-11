package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Elems := make(map[int]int, len(arr2))
	for _, elem := range arr2 {
		arr2Elems[elem] = 0
	}
	needToSortElems := make([]int, 0, len(arr1))

	for _, elem := range arr1 {
		if _, ok := arr2Elems[elem]; ok {
			arr2Elems[elem]++
		} else {
			needToSortElems = append(needToSortElems, elem)
		}
	}

	result := make([]int, 0, len(arr1))
	for _, elem := range arr2 {
		for i := 0; i < arr2Elems[elem]; i++ {
			result = append(result, elem)
		}
	}

	sort.Ints(needToSortElems)

	result = append(result, needToSortElems...)

	return result
}

func main() {
	// [2,2,2,1,4,3,3,9,6,7,19]
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))

	// [22,28,8,6,17,44]
	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
