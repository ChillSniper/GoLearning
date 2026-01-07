package main

import "testing"

func TestSeveral(t *testing.T) {
	myReader := new(MyWriterReader)
	myReader.MyReader()

	myWriter := MyWriterReader{}
	myWriter.MyWriter("Jack")
}
