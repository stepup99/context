package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()
 
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchApi(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}

func fetchApi(ctx context.Context, url string, result chan string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error requesting for %v , err - %v", url, err.Error())
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error making request to %s %s", url, err.Error())
		return
	}

	defer resp.Body.Close()
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		result <- fmt.Sprintf("Error decoding JSON from %s , err %v", url, err)
		return
	}

	result <- fmt.Sprintf("statusCode %v , result %v,", resp.StatusCode, data)
}
