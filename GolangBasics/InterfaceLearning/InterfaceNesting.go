package main

import "fmt"

type A interface {
	runA()
}

type B interface {
	runB()
}

type C interface {
	A
	B
	runC()
}

type Runner struct {
}

func (r Runner) runA() {
	fmt.Println("A is on !")
}

func (r Runner) runB() {
	fmt.Println("B is on !")
}

func (r Runner) runC() {
	fmt.Println("C is on !")
}
