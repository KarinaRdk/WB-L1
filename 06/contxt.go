//  Останавиваем работу горутины через дедлайн контекста

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)

	go print(ctx, &wg)
	wg.Wait()
}

func print(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping...")
			return
		default:
			fmt.Println("Printing...")
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}
}
