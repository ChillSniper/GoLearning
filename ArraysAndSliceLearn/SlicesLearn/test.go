package SlicesLearn

import (
	"fmt"
)

func Test() {
	fruits := []string{"apple", "orange", "grapefruit"}
	fruits = append(fruits, fruits[0])
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
}
