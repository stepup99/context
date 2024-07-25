package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "USERID", 22)

	go perFormTask(ctx)
	time.Sleep(2 * time.Second)
}

func perFormTask(ctx context.Context) {
	userID := ctx.Value("USERID")
	fmt.Println("userID -> ", userID)
}
