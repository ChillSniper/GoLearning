package main

import (
	"sync"
	"testing"
)

func TestMyGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go MyGoroutine("goRoutineA", &wg)
	go MyGoroutine("goRoutineB", &wg)

	wg.Wait()
}
