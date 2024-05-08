package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	squareArr(a)
}

// squareArr принимает массив, итерируется по нему и возводит каждый элемент в квадрат
// каждый элемент обрабатывается в отдельной горутине.
// Выполнение всех рутин обеспечивается использованием счетчика WaitGroup
func squareArr(a [5]int) {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for _, i := range a {
		square(i, &wg)
	}
	wg.Wait()
}

// square принимает int, выводит в stdout его квадрат и уменьшает значение счетчика WaitGroup
// на 1 перед выходом из функции
func square(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i * i)
}
