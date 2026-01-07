package main

import "testing"

func TestNesting(t *testing.T) {
	runner := new(Runner)
	runner.runA()
	runner.runB()
	runner.runC()
}
