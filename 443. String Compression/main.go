package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	toInsert := 0
	c := chars[0]
	count := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == c {
			count++
			continue
		}

		chars[toInsert] = c
		toInsert++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for _, countcountStr := range []byte(countStr) {
				chars[toInsert] = countcountStr
				toInsert++
			}
		}

		count = 1
		c = chars[i]
	}
	chars[toInsert] = c
	toInsert++

	if count > 1 {
		countStr := strconv.Itoa(count)
		for _, cByte := range countStr {
			chars[toInsert] = byte(cByte)
			toInsert++
		}
	}

	chars = chars[:toInsert]

	return toInsert
}

func main() {
	// ['a','2','b','2','c','3']
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	l := compress(chars)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		fmt.Printf("%c", chars[i])
	}
	fmt.Println()

	// ['a','b','1','2']
	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	l = compress(chars)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		fmt.Printf("%c", chars[i])
	}
	fmt.Println()

	// ['a']
	chars = []byte{'a'}
	l = compress(chars)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		fmt.Printf("%c", chars[i])
	}
	fmt.Println()
}
