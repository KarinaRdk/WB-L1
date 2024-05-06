/*
Реализовать быструю сортировку массива (qSort) встроенными методами языка
Среднее время работы O(nlogn), что является асимптотически оптимальным временем работы для алгоритма, основанного на сравнении.
Хотя время работы алгоритма для массива из n элементов в худшем случае может составить Θ(n2), на практике этот алгоритм является одним из самых быстрых.

*/
package main

import (
	"fmt"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i
}

func quickSort(arr []int, low, high int) {
	if len(arr) < 2 {
		return
	}
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex)
		quickSort(arr, pivotIndex+1, high)
	}
}

func main() {
	arr := []int{10, 9, 3, 8, 2, 1, 0, 7, 6, 5, 4}
	a := []int{}
	b := []int{5, 3}
	quickSort(arr, 0, len(arr)-1)
	quickSort(a, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println(a)
	fmt.Println(b)

}
