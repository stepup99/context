package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performCtx(ctx)
	for {
		fmt.Println(time.Now())
		select {
		case <-ctx.Done():
			fmt.Println("Task TimeOut")
			return
		}
	}

}

func performCtx(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("task completed successfully")
	}
}
