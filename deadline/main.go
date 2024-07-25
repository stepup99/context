package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	go performTask(ctx)
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("meet the deadline", ctx.Err())
		return
	}
}
