package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	fmt.Println(sumSquares(a))
}

// sumSquares принимает массив, итерируется по нему, возводит каждый элемент в квадрат
// в отдельной горутине и считает сумму всех квадратов
// Выполнение всех рутин обеспечивается использованием счетчика WaitGroup
func sumSquares(a [5]int) int {
	wg := sync.WaitGroup{}
	sum := 0
	wg.Add(5)
	for _, i := range a {
		sum += square(i, &wg)
	}
	wg.Wait()
	return sum
}

// square принимает int, выводит в stdout его квадрат и уменьшает значение счетчика WaitGroup
// на 1 перед выходом из функции
func square(i int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return i * i
}
