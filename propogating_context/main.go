package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "USERID", 123)
	go performTask(ctx)
	time.Sleep(2 * time.Second)
}

func performTask(ctx context.Context) {
	userID := ctx.Value("USERID")
	fmt.Println("User ID -> ", userID)
}
