package DeferLearning

import "fmt"

func DeferRun() {
	num := 1
	defer fmt.Printf("num is %d\n", num)
	num = 2
	fmt.Printf("num is %d\n", num)
	return
}

func DeferPrintArrTest() {
	arr := []int{1, 2, 3, 4, 5, 6}
	defer printArr(&arr)
	arr[0] = 100
	return
}

func printArr(arr *[]int) {
	for i := range *arr {
		fmt.Println((*arr)[i])
	}
}

func OperateFinalResult() int {
	num := 1

	defer func() {
		num++
	}()

	return num
}
