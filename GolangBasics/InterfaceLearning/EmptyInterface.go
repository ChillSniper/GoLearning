package main

import "fmt"

type Student struct {
	Name   string
	id     int
	height int
	weight float64
}

func TestEmpty() {
	var any interface{}
	any = 10
	fmt.Println(any)

	any = "10"
	fmt.Println(any)

	any = Student{
		Name:   "jack",
		id:     10,
		height: 20,
		weight: 3.0,
	}
	fmt.Println(any)
}
