// Реализовать все возможные способы остановки выполнения горутины.
//  ЗЧитаем в рутине из канала, пока он открыт

package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go print(ch, &wg)
	go send(ch)
	wg.Wait()
}

func send(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch) // Отправляем сигнал о остановке
}

func print(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//  будет получать данные пока канал не закрыт
	for d := range ch {
		fmt.Println(d)
	}
}
