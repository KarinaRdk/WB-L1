// Реализовать все возможные способы остановки выполнения горутины.
//  Канал уведомлений

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//  Создаем канал для отправки уведомлений без данных
	sygnalChan := make(chan struct{})

	go print(sygnalChan, &wg)
	time.Sleep(3 * time.Second)
	// Отправляем сигнал о остановке
	sygnalChan <- struct{}{}
	wg.Wait()
}

func print(sygnalChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-sygnalChan:
			fmt.Println("Stopping after sygnal was received")
			return
		default:
			fmt.Println("Printing...")
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}
}
