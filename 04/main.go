package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func workers(c chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case i, ok := <-c:
			if!ok {
				return
			}
			fmt.Println(i)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numberOfWorkers := 3

	c := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go workers(c, &wg, ctx)
	}

	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				c <- i
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Listen for OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs

	// Cancel the context to signal the producer goroutine to stop
	cancel()

	wg.Wait()
}
