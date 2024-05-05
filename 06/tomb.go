package main

import (
	"fmt"
	"github.com/tidwall/tomb"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	tomb := tomb.New()

	go print(tomb, &wg)
	wg.Wait()
	tomb.Die(nil) // Останавливаем горутину
}

func print(tomb *tomb.Tomb, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-tomb.Dying():
			fmt.Println("Stopping...")
			return
		default:
			fmt.Println("Printing...")
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}
}
