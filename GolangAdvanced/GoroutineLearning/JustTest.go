package main

import (
	"fmt"
	"time"
)

func myGoroutine() {
	fmt.Println("my goroutine")
}

func main() {
	go myGoroutine()
	fmt.Println("end the program")
	time.Sleep(1 * time.Second)
}
