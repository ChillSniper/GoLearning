package PanicLearning

import "fmt"

func RecoverExceptionHandling() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现panic，使用recover获取信息：", error)
		}
	}()
	fmt.Print("before")
	panic("appear panic")
	fmt.Print("after")
}
