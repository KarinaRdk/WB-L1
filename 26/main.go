package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcd"
	b := "abCdefAaf"
	c := "aabcd"
	fmt.Println(a, ifUnique(a))
	fmt.Println(b, ifUnique(b))
	fmt.Println(c, ifUnique(c))
}

// ifUnique принимает строку, итерируется по всем символам в ней и проверяет, встречался
// ли этот символ в строке ранее, используя map, возвращает true, если все символы уникальны,
// нет - в обратном случае, во входящей строке все Unicode символы приводятся к нижнему регистру
// для обеспечения регистронезависимости
func ifUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]struct{})
	for _, i := range s {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
			continue
		}
		return false
	}
	return true
}
