package ArraysLearn

import "fmt"

var Scores [3]int
var prices = [3]float64{19, 9, 1.1}

func Test() {
	names := [...]string{"张三", "李四", "王五"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
