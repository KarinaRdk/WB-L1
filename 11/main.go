// Реализовать пересечение двух неупорядоченных множеств

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 26, 38, 4}
	fmt.Println(f(a, b)) // Вывод: [1 4]
	a = []int{}
	b = []int{1, 26, 38, 4}
	fmt.Println(f(a, b)) // Вывод: []
}

func f(a []int, b []int) []int {
	// Если один из срезов пустой - можем вернуть ответ сразу
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	// Сортируем оба среза
	sort.Ints(a)
	sort.Ints(b)

	// Инициализируем указатели для обоих срезов
	i, j := 0, 0

	// Инициализируем срез для хранения общих элементов
	var commonElements []int

	// Проходимся по обоим срезам
	for i < len(a) && j < len(b) {
		// Если текущие элементы обоих срезов равны
		if a[i] == b[j] {
			// Добавляем в общий срез и двигаем оба указателя вперед
			commonElements = append(commonElements, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			// Если элемент а меньше, двигаем 'i' вперед
			i++
		} else {
			// Если элемент b меньше, двигаем 'j' вперед
			j++
		}
	}

	return commonElements
}
