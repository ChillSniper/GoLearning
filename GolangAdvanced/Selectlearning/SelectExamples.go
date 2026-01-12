package Selectlearning

import (
	"fmt"
	"time"
)

func LearnNoDefaultAndCaseError() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	select {
	case <-ch1:
		fmt.Printf("receive ch1\n")
	case num := <-ch2:
		fmt.Printf("receive ch2: %d\n", num)
	}

}

func SeveralCaseAndDefault() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 3; i++ {
			select {
			case v := <-ch1:
				fmt.Printf("receive ch1, val = %d\n", v)
			case v := <-ch2:
				fmt.Printf("receive ch2, val = %d\n", v)
			default:
				fmt.Printf("default ...")
			}
			time.Sleep(1 * time.Second)
		}
	}()

	ch1 <- 1
	time.Sleep(1 * time.Second)
	ch2 <- 2
	time.Sleep(4 * time.Second)
}

func RandomChoice() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 5
	ch2 <- 3
	select {
	case v := <-ch1:
		fmt.Printf("receive ch1, val = %d\n", v)
	case v := <-ch2:
		fmt.Printf("receive ch2, val = %d\n", v)
	default:
		fmt.Printf("default ...")
	}
}
