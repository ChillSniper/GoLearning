package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Printf("a = %s, b = %d, c = %d\n", a, b, c)
}
