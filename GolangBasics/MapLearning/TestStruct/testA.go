package main

import (
	"fmt"
	"goStudy/GolangBasics/structLearning"
)

func main() {
	st := structLearning.Student{
		ID:    100,
		Name:  "ZhangSan",
		Age:   18,
		Score: 98,
	}
	fmt.Println(st)
}
