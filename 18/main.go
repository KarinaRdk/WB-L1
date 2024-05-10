// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Структура counter представляет собой счетчик, который может быть инкрементирован в конкурентной среде.
type counter struct {
	v  int64        // v - текущее значение счетчика
	mu sync.Mutex // mu - блокировка для синхронизации доступа к счетчику
}

func main() {
	withMu()
	withAtomic()
	withChan()
}

//  Имплементация с использованием мьютекса
func withMu() {
	c := counter{v: 0}     // Создаем экземпляр структуры counter с начальным значением 0
	wg := sync.WaitGroup{} // Создаем WaitGroup для ожидания завершения горутин
	wg.Add(70)             // Добавляем 20 задач в WaitGroup

	for _ = range 70 {
		go func() {  //  Создаем 20 горутин
			defer wg.Done()  //  Уменьшаем счетчик
			c.mu.Lock()  //  Лочим мьютекс и блокируем другие рутины, которые пытаются получить доступ к этой же переменной
			c.v++  //  Меняем значение
			c.mu.Unlock()  //  Снимаем блок для других горутин
		}()

	}

	wg.Wait()        // Ожидаем завершения всех горутин
	fmt.Println("mutex: ", c.v) // Выводим итоговое значение счетчика
}
 //  Используем атомик, чтобы обеспечить атомарную инкрементацию счетчика

func withAtomic() {
c := counter{v: 0} 
var wg sync.WaitGroup
wg.Add(70)
for _= range 70 {
	go func() {
		defer wg.Done()
		atomic.AddInt64(&c.v, 1)
	}()
	
}
wg.Wait()
fmt.Println("atomic: ", c.v)
}

//  Используем канал, чтобы заблокировать другие рутины на время работы с переменной 
func withChan() {
	c := counter{v:0}
	var wg sync.WaitGroup
	wg.Add(70)
	ch := make(chan struct{}, 1)

	for i := 0; i < 70; i++ {
		go func(){
			defer wg.Done()
			ch <-struct{}{}
			c.v++
			<-ch
		}()
	}
	wg.Wait()
	fmt.Println("chanel: ", c.v)
}