package main

import "fmt"

type MyWriter interface {
	MyWriter(s string)
}

type MyReader interface {
	MyReader()
}

type MyWriterReader struct {
}

func (r MyWriterReader) MyWriter(s string) {
	fmt.Printf("call MyWriterReader's MyWriter %s\n", s)
}

func (r MyWriterReader) MyReader() {
	fmt.Printf("call MyWriterReader's MyReader\n")
}
