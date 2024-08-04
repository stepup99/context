package main

import (
	"context"
	"fmt"
)

func main() {
	// declare the context object and pass in a value
	// ctx := context.WithValue(parentContext , key, value)
	// access values stored within context
	// val := ctx.Value(key)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "123")
	data(ctx)
}

func data(ctx context.Context) {
	key := ctx.Value("key")
	if key != nil {
		fmt.Printf("key -> %v", key)
	} else {
		fmt.Println("key not found")
	}
}

// WithCancel: returns a new context and cancel function
// WithTimeout: returns a new context with timeout duration
