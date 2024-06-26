/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

package main

import (
	"fmt"
)

func main() {
	t := []float32{25.4, -27.0, 13.0, 19.0, 0.0, 5.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(f(t))

}

// Функция объединит в подмножества, определив количество десятков в числе и используя эту величину как ключ в map,
// где и будут храниться все значения
func f(t []float32) map[int][]float32 {
	m := make(map[int][]float32)
	for i, v := range t {
		key := int(v) / 10
		fmt.Println(key)
		m[key] = append(m[key], t[i])
	}
	return m
}
