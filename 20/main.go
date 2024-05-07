// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	fmt.Printf("%s.\n", turn(s))
	t := "мама мыла раму"
	fmt.Printf("%s.\n", otherTurn(t))

}

func turn(input string) string {
	// Получаем массив подстрок без пробелов
	substr := strings.Fields(input)
	var buf bytes.Buffer
	for i, v := range substr {
		buf.WriteString(v)
		if i < len(substr)-1 {
			buf.WriteString(" ")
		}

	}
	return buf.String()
}

func otherTurn(input string) string {
	// Получаем массив подстрок без пробелов
	substr := strings.Fields(input)
	j := len(substr) - 1
	for i := range substr {
		if i > (len(substr)-1)/2 {
			break
		}
		substr[i], substr[j] = substr[j], substr[i]
		j--
	}

	return strings.Join(substr, " ")
}
