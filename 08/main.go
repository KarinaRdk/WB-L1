// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
	"fmt"
)

func main() {
	var num int64
	fmt.Printf("Enter number:")
	fmt.Scanf("%d", &num)
	fmt.Printf("Number %d in binary is %b\n", num, num)

	var pos int64
	fmt.Println("Position of bit: ")
	fmt.Scan(&pos)

	var bitValue int64
	fmt.Println("Value of bit: ")
	fmt.Scan(&bitValue)

	if (pos < 0 || pos > 63) || (bitValue != 0 && bitValue != 1) {
		fmt.Println("wrong position or value of bit")
		return
	}

	result := setBit(num, pos, bitValue)
	fmt.Printf("Number %d in binary is %b\n", result, result)
}

// setBit принимает число, позицию бита (индекс бита в двоичном представлении числа,
//	начиная с 0 справа налево), желаемое значение бита и меняет число в
//  оответствии с этими параметрами

func setBit(num int64, pos int64, bitValue int64) int64 {
	// Сператор |= выполняет побитовую операцию ИЛИ с
	// текущим значением num и результатом сдвига bitValue << pos,
	// тем самым устанавливая бит на позиции pos в num.
	shifted := bitValue << pos
	num |= shifted
	return num
}
