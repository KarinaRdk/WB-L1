// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 2
	b := 3
	//  go позволяет менять значения переменных очень просто:
	a, b = b, a
	fmt.Println(a, b)
	//  На случай, если авторы задания хотели что-то сложнее:
	swap(a, b)
	aORSwap(a, b)

}

func swap(a, b int) {
	fmt.Printf("До обмена числа: %d и %d\n", a, b)
	b = a + b
	a = b - a
	b = b - a
	fmt.Printf("После обмена числа: %d и %d\n", a, b)
}

func aORSwap(a, b int) {
	fmt.Printf("До обмена числа: %d и %d\n", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("После обмена числа: %d и %d\n", a, b)
}
