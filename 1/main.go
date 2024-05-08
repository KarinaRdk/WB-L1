package main

import "fmt"

// поле структуры является встреным, если оно объявлено без имени
// встраивание используется для продвижения полей и методов встроенной структуры
// к встроенным полям есть доступ у внешних потребителей

type Human struct {
	walk string
	talk string
}

type Action struct {
	Human
}

func main() {
	Steave := Human{walk: "swiftly", talk: "loud"}
	Steave.toHuman()
	a := Action{Steave}
	a.toHuman()
}

func (h Human) toHuman() {
	fmt.Println(h.talk)
}
