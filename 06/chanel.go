// Реализовать все возможные способы остановки выполнения горутины.
//  Закрываем канал и обуславливаем этим выход из цикла

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	stopChan := make(chan bool)

	go print(stopChan, &wg)
	time.Sleep(3 * time.Second)
	close(stopChan) // Отправляем сигнал о остановке
	wg.Wait()
}

func print(stopChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Stopping after a chanel was closed")
			return
		default:
			fmt.Println("Printing...")
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}
}
