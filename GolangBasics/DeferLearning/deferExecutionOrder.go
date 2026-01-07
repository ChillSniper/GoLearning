package DeferLearning

import "fmt"

func deferA() {
	fmt.Println("deferA")
}

func deferB() {
	fmt.Println("deferB")
}

func deferC() {
	fmt.Println("deferC")
}

func Ts() {
	defer deferA()
	defer deferB()
	defer deferC()
}
