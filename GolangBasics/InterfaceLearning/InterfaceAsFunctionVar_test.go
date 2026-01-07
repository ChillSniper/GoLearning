package main

import (
	//"fmt"
	"testing"
)

func TestDoJob(t *testing.T) {
	myReaderNew := &MyReaderNew{1, 3}
	DoJob(myReaderNew)
}
