package SyncLearning

import (
	"fmt"
	"sync"
	"time"
)

var cnt = 0

func DeadLock() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	copyMutex(mu)
}

func copyMutex(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("ok")
}

func CircularDependency() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(1 * time.Second)

		mu2.Lock()

		defer mu2.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()

		time.Sleep(1 * time.Second)
		mu1.Lock()
		defer mu1.Unlock()
	}()

	wg.Wait()
}

var m = make(map[string]int)

func LearnMap() {
	wg := sync.WaitGroup{}
	n := 10000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(num int) {
			defer func() {
				defer wg.Done()
				mu.Unlock()
			}()
			key := fmt.Sprintf("key-%d", num)
			mu.Lock()
			setVal(key, num)
			fmt.Printf("key:=%s, val=%d\n", key, getVal(key))
		}(i)
	}

	wg.Wait()
}

func getVal(key string) int {
	return m[key]
}

func setVal(key string, val int) {
	m[key] = val
}

func LearnSyncMap() {
	m := sync.Map{}

	m.Store("name", "zhangsan")
	m.Store("age", 18)

	age, _ := m.Load("age")
	fmt.Println(age.(int))

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key is %s, val is %d\n", key, value)
		return true
	})

	m.Delete("age")
	age, ok := m.Load("age")
	fmt.Println(ok, age)

	m.LoadOrStore("name", "zhangsan")
	name, _ := m.Load("name")
	fmt.Println(name)
}
