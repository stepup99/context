package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		// operation goes here
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("completed in 5 second")
		case <-ctx.Done():
			fmt.Println("context done ------->")
		}
	}(ctx)
	cancel()
	time.Sleep(7 * time.Second)
}
