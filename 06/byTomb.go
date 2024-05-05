package main

import (
	"fmt"
	"launchpad.net/tomb"
	"time"
)

type Proc struct {
	Tomb tomb.Tomb
}

func (proc *Proc) Exec() {
	defer proc.Tomb.Done()
	for {
		select {
		case <-proc.Tomb.Dying():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Printing")
		}
	}
}

func main() {
	proc := &Proc{}
	go proc.Exec()
	time.Sleep(1 * time.Second)
	proc.Tomb.Kill(fmt.Errorf("Stoped"))
	err := proc.Tomb.Wait() // Will return the error that killed the proc
	fmt.Println(err)
}
