package main

import "fmt"

type Reader interface {
	Read() int
}

type MyReaderNew struct {
	a, b int
}

func (m *MyReaderNew) Read() int {
	return m.a * m.b
}

func DoJob(r Reader) {
	fmt.Printf("myReaderNew is %d\n", r.Read())
}
