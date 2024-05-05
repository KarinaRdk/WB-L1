package main

import (
	"fmt"
	"os"
	"strconv"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func startWork(duration time.Duration, ch chan int, wg *sync.WaitGroup) {
	// создаем таймаут заданной пользователем пролдолжительности
	timeout := time.After(duration)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-timeout: //  проверяем, не истекло ли время
				fmt.Println("Время истекло")
				os.Exit(0)
			case data := <-ch:
				fmt.Println("Данные получены:", data)
			}
		}
	}()
}

// GracefulShutdown обеспечивает корректное завершение работы программы.
func GracefulShutdown(ch chan int, wg *sync.WaitGroup) {
	interruptChannel := make(chan os.Signal, 1) // Создаем канал для приема сигналов прерывания
	signal.Notify(interruptChannel, syscall.SIGINT, syscall.SIGTERM) // Уведомляем, что хотим получать сигналы SIGINT и SIGTERM
	<-interruptChannel // Ждем сигнал прерывания

	close(ch) // Закрываем канал данных
	wg.Wait() // Ждем завершения всех рабочих

	fmt.Println("\nВыход выполнен") // Выводим сообщение о завершении работы
}

func usage() {
	fmt.Println("Usage: go run main.go <duration>")
	os.Exit(1)
}

func main() {
// проверяем количество аргументов объясняем пользователю, что он дал не правильный ввод
	if len(os.Args) < 2 {
		usage()
	}

	a, err := strconv.Atoi(os.Args[1])
	if err!= nil {
		usage()
	}
	
	duration := time.Duration(a) * time.Second

	ch := make(chan int)
	// Отправляем данные в канал
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Second)
		}
	}()
	//  Создаем счетчик горутин, чтобы программа не завершилась раньше времени
	var wg sync.WaitGroup
	wg.Add(1)
	startWork(duration, ch, &wg)
	GracefulShutdown(ch, &wg)
}
