package main

import (
	"fmt"
	"log"
	"sync"
)

type InMemory struct {
	m map[int]int
	//  A read/write mutex позволяет всем читателям одновременно получить доступ к map, но писатель заблокирует всех остальных
	lock sync.RWMutex
}

// New инициализирует новую карту для хранения map
func New() *InMemory {
	m := make(map[int]int)
	c := InMemory{m: m}
	return &c
}

// Set добавляет новый элемент в map
func (c *InMemory) Set(key int, value int) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.m[key]; !ok {
		c.m[key] = value
		log.Print("Добавлено в map ", key, string(value))

		return nil
	}
	return fmt.Errorf("Уже существует в map")
}

// Get считывает данные из map и возвращает найденное значение для предоставленного ключа и ошибку
func (c *InMemory) Get(key int) (value int, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if mapValue, ok := c.m[key]; ok {
		return mapValue, nil
	}

	return value, fmt.Errorf("Такого значения нет")
}

func main() {
	m := New()
	m.Set(1, 3)
	value, err := m.Get(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
