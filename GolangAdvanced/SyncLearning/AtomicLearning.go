package SyncLearning

import (
	"fmt"
	"sync/atomic"
)

type Student struct {
	Name string
	Age  int
}

func LearnAtomicUsage() {
	st1 := Student{
		Name: "ZhangSan",
		Age:  18,
	}

	st2 := Student{
		Name: "JinChuan",
		Age:  19,
	}

	//st3 := Student {
	//	Name: "Jobin",
	//	Age : 20,
	//}

	var v = atomic.Value{}
	v.Store(st1)

	fmt.Println(v.Load().(Student))

	old := v.Swap(st2)
	fmt.Printf("after swap: v = %v\n", v.Load().(Student))
	fmt.Printf("after swap: old= %v\n", old)
}
