package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	go func(ctx context.Context) {

		select {
		case <-time.After(10 * time.Second):
			fmt.Println("i am here 1")
			fmt.Println("after 10 second")
		case <-ctx.Done():
			fmt.Println("i am here 2")
			fmt.Println("this is done ")
		case <-time.After(12 * time.Second):
			fmt.Println("i am here 3")
			fmt.Println("this is done ")
		default:
			fmt.Println("this is default")
		}
		fmt.Println("i am here")
		// deadtime, ok := ctx.Deadline()
		// if ok == true {
		// 	fmt.Println("this is true ", deadtime)
		// }
		

	}(ctx)

	cancel()
	time.Sleep(10 * time.Second)

}
