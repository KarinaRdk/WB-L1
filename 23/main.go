package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("before: ", a)
	a, _ = deleteFromSlice(a, 1)
	fmt.Println("after:  ", a)
	a, _ = deleteFromSlice(a, 4)
	fmt.Println("after:  ", a)

}

// deleteFromSlice принимает срез интов и значение индекса и возвращает новый срез
// без i-того элемента и ошибку, первоначально переданный срез остается неизменным
func deleteFromSlice(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, fmt.Errorf("index is out of range")
	}
	c := make([]int, i)
	copy(c, s[:i])
	if i+1 < len(s) {
		c = append(c, s[i+1:]...)
	}
	return c, nil
}
