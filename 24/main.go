/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.

Формула вычисления расстояния между двумя точками A(xa, ya) и B(xb, yb) на плоскости:
AB = √((xb - xa)2 + (yb - ya)2)
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func main() {
	//  Создаем переменные для хранения чисел, с которыми будем производить вычисления
	var x1, y1, x2, y2 float64

	//  Получаем значения от пользователя
	fmt.Print("Введите координаты x y первой точки через пробел: ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("Введите координаты x y второй точки через пробел: ")
	fmt.Scanln(&x2, &y2)
	//  Создаем точки
	a := newPoint(x1, y1)
	b := newPoint(x2, y2)
	fmt.Println("Расстояние между точками равно: ", distance(a, b))
}
