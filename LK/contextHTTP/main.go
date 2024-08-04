package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Associate the context with the HTTP request
	req = req.WithContext(ctx)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
