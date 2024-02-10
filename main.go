package main

import (
	"context"
	"time"
)

const timeout = time.Second * 2

func main() {
	parent := context.Background()
	realMain(parent, 0)
	time.Sleep(3 * time.Second)
}

func realMain(ctx context.Context, num int) {
	cont, _ := context.WithTimeout(ctx, timeout)
	if num == 0 {
		return
	}
	for i := 0; i <= num; i++ {
		go worker(cont)
	}
}

func worker(ctx context.Context) {
	<-ctx.Done()
	println("worker done")
}
