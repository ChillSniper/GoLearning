package main

import (
	"fmt"
	"sync"
	"time"
)

func MyGoroutine(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("myGoroutine %s\n", name)
		time.Sleep(10 * time.Millisecond)
	}

}
