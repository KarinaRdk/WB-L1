// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

// Структура counter представляет собой счетчик, который может быть инкрементирован в конкурентной среде.
type counter struct {
	v  int        // v - текущее значение счетчика
	mu sync.Mutex // mu - блокировка для синхронизации доступа к счетчику
}

func main() {
	c := counter{v: 0}     // Создаем экземпляр структуры counter с начальным значением 0
	wg := sync.WaitGroup{} // Создаем WaitGroup для ожидания завершения горутин
	wg.Add(2)              // Добавляем 2 задачи в WaitGroup

	// Запускаем две горутины, каждая из которых будет инкрементировать значение счетчика
	go write(&c, &wg)
	go writeTwo(&c, &wg)

	wg.Wait()        // Ожидаем завершения всех горутин
	fmt.Println(c.v) // Выводим итоговое значение счетчика
}

// write - функция, которая инкрементирует значение счетчика в конкурентной среде
func write(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()     // Указываем, что горутина завершена
	c.mu.Lock()         // Блокируем доступ к счетчику
	defer c.mu.Unlock() // Освобождаем блокировку после завершения функции
	for i := 0; i < 5; i++ {
		c.v += i
	}
}

// writeTwo - функция, которая также инкрементирует значение счетчика в конкурентной среде
func writeTwo(c *counter, wg *sync.WaitGroup) {
	defer wg.Done()     // Указываем, что горутина завершена
	c.mu.Lock()         // Блокируем доступ к счетчику
	defer c.mu.Unlock() // Освобождаем блокировку после завершения функции
	for i := 0; i < 6; {
		c.v += i
		i += 2
	}
}
