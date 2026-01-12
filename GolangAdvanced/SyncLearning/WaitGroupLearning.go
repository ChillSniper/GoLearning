package SyncLearning

import (
	"fmt"
	"sync"
)

func UsingChannelsToSync() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("Num:%d\n", i)
			ch <- i
		}(i)

	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("end")
}

var wg sync.WaitGroup

func MyGoroutine() {
	T := func() {
		defer wg.Done()
		fmt.Println("myGoroutine!")
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go T()
	}
	wg.Wait()
	fmt.Println("end")
}
