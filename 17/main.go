// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func binarySearch(s []int, n int) int {
	lowerBound := 0
	upperBound := len(s)
	for upperBound > lowerBound {

		if n > s[upperBound/2] {
			lowerBound = upperBound / 2
		}
		if n < s[upperBound/2] {
			upperBound = upperBound / 2

		}
		if n == s[upperBound/2] {
			return upperBound / 2
		}
	}

	return -1
}

func main() {
	test()
}

func test() {
	s := []int{1, 2, 3, 4, 5}
	a := binarySearch(s, 0)
	if a != -1 {
		fmt.Println("a wrong!, expected -1, got ", a)
	} else {
		fmt.Println("OK")
	}

	s1 := []int{-5, 0, 10, 11}
	c := binarySearch(s1, 0)
	if c != 1 {
		fmt.Println("c wrong!, expected 1, got ", c)
	} else {
		fmt.Println("OK")
	}

	s2 := []int{}
	b := binarySearch(s2, 3)
	if b != -1 {
		fmt.Println("b wrong!, expected -1, got ", b)
	} else {
		fmt.Println("OK")
	}

}
