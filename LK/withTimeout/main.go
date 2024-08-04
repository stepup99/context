package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("after 5 second")
		case <-ctx.Done():
			fmt.Println("ctx done >>>>>>> ")
		}
	}(ctx)
	cancel()
	time.Sleep(8 * time.Second)
}
