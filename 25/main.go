// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func Sleep(milliseconds time.Duration) {
	<-time.After(milliseconds * time.Millisecond)
}

func main() {
	fmt.Println("Start")

	Sleep(2000) 

	fmt.Println("End")
}