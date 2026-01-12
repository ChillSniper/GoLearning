package ContextLearning

import (
	"context"
	"fmt"
	"time"
)

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit !\n", name)
			return
		default:
			fmt.Printf("%s watching...\n", name)
			time.Sleep(time.Second)
		}
	}
}

func T_Cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go Watch(ctx, "goRoutineA")
	go Watch(ctx, "goRoutineB")

	time.Sleep(6 * time.Second)
	fmt.Println("end Watching !!!")

	cancel()
	time.Sleep(time.Second * 2)
}

func T_Deadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))
	defer cancel()
	go Watch(ctx, "goRoutineA")
	go Watch(ctx, "goRoutineB")

	time.Sleep(6 * time.Second)
	fmt.Println("end Watching !!!")
}

func print_func(ctx context.Context) {
	fmt.Printf("name is %s\n", ctx.Value("name").(string))
}

func T_WithValue() {
	ctx := context.WithValue(context.Background(), "name", "Herbert Lu")
	go print_func(ctx)
	time.Sleep(1 * time.Second)
}
