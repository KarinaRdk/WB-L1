// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func main() {
	s := "12345"
	fmt.Println(turn(s))
	s = "главрыба"
	fmt.Println(turn(s))

}

func turn(input string) string {
	//  поскольку строка - срез байтов, и разные символы UNICOD занимают разное боличество байтов, нам придетя привести строку к срезу рун
	s := []rune(input)
	j := len(s) - 1
	//  проходим по срезу до середины и меняем руны местами
	for i := range s {
		if i == (len(s)-1)/2 {
			break
		}
		s[i], s[j] = s[j], s[i]
		j--
	}

	return string(s)
}
