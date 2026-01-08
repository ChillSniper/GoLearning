package ChannelLearning

import "fmt"

func Init() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	v := <-ch
	fmt.Println("v:", v)
	close(ch)
}
