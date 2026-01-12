package SyncLearning

import (
	"fmt"
	"sync"
)

var (
	num int
	//wg  = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func add() {
	mu.Lock()
	defer wg.Done()
	num++
	mu.Unlock()
}

func Action() {
	var n = 10 * 10 * 10 * 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go add()
	}
	wg.Wait()

	fmt.Println(num == n)
	fmt.Printf("num = %d, n = %d\n", num, n)
}
