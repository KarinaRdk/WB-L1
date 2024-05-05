package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"strconv"
)

// Worker представляет собой структуру рабочего, который обрабатывает данные из канала.
type Worker struct {
	wg *sync.WaitGroup
	ch chan int
}

// workerPool - это тип, представляющий пул рабочих.
type workerPool []Worker

// CreatePool создает пул рабочих задач с указанным количеством рабочих.
func CreatePool(n int, wgr *sync.WaitGroup, cha chan int) workerPool {
	workerPool := make(workerPool, 0, 5) 
	for i := 0; i < n; i++ { 
		workerPool = append(workerPool, Worker{wg: wgr, ch: cha}) 
	}
	return workerPool 
}

// Work метод выполняет работу рабочего, обрабатывая данные из канала до его закрытия.
func (w *Worker) Work(index int) {
	defer w.wg.Done() // Откладываем выполнение wg.Done() до завершения функции
	for data := range w.ch { // Обрабатываем данные из канала до его закрытия
		fmt.Printf("Рабочий %d получил данные - %d\n", index, data) // Выводим информацию о полученных данных
	}
}

// StartWork запускает все рабочие в пуле в отдельных рутинах.
func (p *workerPool) StartWork() {
	for i, worker := range *p { 
		go worker.Work(i) 
	}
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

func main() {
	if len(os.Args) < 2 { // Проверяем, переданы ли аргументы командной строки
		usage() // Если нет, объясяем польователю, какой ввод нужен
	}
	a, err := strconv.Atoi(os.Args[1]) // Преобразуем первый аргумент командной строки в число
	if err!= nil { // Если произошла ошибка преобразования
		usage() // Выводим сообщение об использовании и завершаем программу
	}
	var wg sync.WaitGroup // Создаем WaitGroup для ожидания завершения всех рабочих
	wg.Add(a) // Увеличиваем счетчик WaitGroup на количество рабочих

	ch := make(chan int) // Создаем канал для передачи данных

	p := CreatePool(a, &wg, ch) // Создаем пул рабочих

	p.StartWork() // Запускаем рабочих

	go func() { // Запускаем горутину для генерации данных в бесконечном цикле
		for i := 0; ; i++ { 
			ch <- i
			time.Sleep(time.Second) 
		}
	}()

	GracefulShutdown(ch, &wg) // Вызываем функцию для корректного завершения работы программы
}

// usage выводит сообщение об использовании и завершает программу
func usage() {
	fmt.Println("Использование: go run 04/main.go <количество рабочих>")
	os.Exit(1)
}
