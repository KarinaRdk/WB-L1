/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

- Неэффективная работа с памятью: будет храниться весть созданный базовый массив, который не очистит GC
хотя необходима только его десятая часть
чтобы это исправить можно работать с копией или создавать срез сразу при вызове функции и не хранить переменну,
которая бы указывала на большой созданный массив и мешала бы сборщику мусора очистить память
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

var justString string

func createHugeString(size int) string {
	// Здесь может быть логика создания большой строки.
	// Для примера, мы просто создаем строку из символов '*'.
	result := ""
	for i := 0; i < size; i++ {
		result += "*"
	}
	return result
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("initially: %d\n", m.Alloc/1024)
	// Функция из задания
	someFunc()
	runtime.ReadMemStats(&m)
	fmt.Printf("after someFunc(): %d\n", m.Alloc/1024)
	runtime.GC()
	time.Sleep(4 * time.Second)
	fmt.Printf("after GC: %d\n", m.Alloc/1024)

	// Создаем подстроку непосредственно при вызове функции.
	substring := createHugeString(1 << 10)[:10]
	fmt.Println(substring)

	runtime.ReadMemStats(&m)
	fmt.Printf("after substring := createHugeString(1 << 10)[:10]: %d\n", m.Alloc/1024)
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("after GC: %d\n", m.Alloc/1024)
	runtime.GC()
	shortstring := make([]byte, 10)

	// Используем копирование
	copy(shortstring, createHugeString(1<<10))
	fmt.Println(string(shortstring))
	runtime.ReadMemStats(&m)
	fmt.Printf("after copy(shortstring, createHugeString(1 << 10)): %d\n", m.Alloc/1024)

	runtime.GC()
	fmt.Printf("after GC: %d\n", m.Alloc/1024)

}
