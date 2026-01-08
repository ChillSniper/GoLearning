package ChannelLearning

import (
	"fmt"
	"time"
)

func Add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func MakeTest() {
	ch := make(chan bool, 1)
	num := 0
	for i := 0; i < 10000000; i++ {
		go Add(ch, &num)
	}
	time.Sleep(time.Second)
	fmt.Println(num)
}
