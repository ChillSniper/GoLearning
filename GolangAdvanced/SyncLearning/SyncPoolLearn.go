package SyncLearning

import (
	"fmt"
	"sync"
)

func LearnSyncPool() {
	pool := sync.Pool{
		New: func() interface{} {
			return &Student{
				Name: "JinChuan",
				Age:  19,
			}
		},
	}

	st := pool.Get().(*Student)

	fmt.Println(st.Name, st.Age)
	fmt.Printf("addr is %p\n", st)

	pool.Put(st)

	stA := pool.Get().(*Student)

	fmt.Println(stA.Name, stA.Age)
	fmt.Printf("addr is %p\n", stA)

}
