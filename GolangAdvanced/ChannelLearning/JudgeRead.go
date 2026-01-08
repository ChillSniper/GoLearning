package ChannelLearning

import (
	"fmt"
	"time"
)

func JudgeRead() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			v, ok := <-ch
			if ok {
				fmt.Printf("v = %d\n", v)
			} else {
				fmt.Printf("channel data has been read over, v = %d\n", v)
			}
		}
	}()
	time.Sleep(1 * time.Second)
}
