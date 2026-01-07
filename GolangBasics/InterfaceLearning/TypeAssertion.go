package main

import "fmt"

func TypeAssertion() {
	var a int = 1
	var i interface{} = a
	b, ok := i.(int)
	fmt.Printf("val is %d, ok is %t", b, ok)
}

func TypeWrong() {
	var x interface{}
	x = "golang"
	val, ok := x.(int)
	fmt.Print(val)
	fmt.Println(ok)
}
