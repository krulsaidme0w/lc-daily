package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	num1Index := len(num1) - 1
	num2Index := len(num2) - 1
	addOne := false
	res := ""

	for num1Index >= 0 && num2Index >= 0 {
		n1 := num1[num1Index]
		n2 := num2[num2Index]

		n1Int := int(n1 - '0')
		n2Int := int(n2 - '0')
		s := n1Int + n2Int
		if addOne {
			s += 1
			addOne = false
		}

		if s >= 10 {
			addOne = true
		}

		res = strconv.Itoa(s%10) + res

		num1Index -= 1
		num2Index -= 1
	}

	for num1Index >= 0 {
		n1 := num1[num1Index]
		n1Int := int(n1 - '0')
		s := n1Int
		if addOne {
			s += 1
			addOne = false
		}

		if s >= 10 {
			addOne = true
		}

		res = strconv.Itoa(s%10) + res
		num1Index -= 1
	}

	for num2Index >= 0 {
		n1 := num2[num2Index]
		n1Int := int(n1 - '0')
		s := n1Int
		if addOne {
			s += 1
			addOne = false
		}

		if s >= 10 {
			addOne = true
		}

		res = strconv.Itoa(s%10) + res
		num2Index -= 1
	}

	if addOne {
		res = "1" + res
	}

	return res
}

func main() {
	// "533"
	num1 := "456"
	num2 := "77"
	fmt.Println(addStrings(num1, num2))

	// "0"
	num1 = "0"
	num2 = "0"
	fmt.Println(addStrings(num1, num2))

	// "10"
	num1 = "9"
	num2 = "1"
	fmt.Println(addStrings(num1, num2))
}
