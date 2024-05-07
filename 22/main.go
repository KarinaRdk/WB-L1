// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {

	//  Создаем переменные для хранения чисел, с которыми будем производить вычисления
	a := new(big.Int)
	b := new(big.Int)
	//  Получаем значения от пользователя
	fmt.Print("Введите первое число: ")
	fmt.Scanln(a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(b)
	// Вызываем функции для расчетов и выыводим их результаты
	fmt.Println("сумма = ", sum(a, b))
	fmt.Println("разность = ", sub(a, b))
	fmt.Println("произведдение = ", mult(a, b))
	val, err := div(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("частное = ", val)
}

func sum(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

func sub(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

func mult(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func div(a, b *big.Int) (*big.Int, error) {
	result := new(big.Int)
	if b.Cmp(big.NewInt(0)) == 0 {
		err := fmt.Errorf("Деление на ноль")
		return result, err
	}
	result.Div(a, b)
	return result, nil
}
