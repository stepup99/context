package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	go httpJson(ctx)
	for {
		fmt.Println("in for loop")
		select {
		case <-ctx.Done():
			fmt.Println("program get intreeputed due to context")
			return
		}

	}

}

func httpJson(ctx context.Context) {
	httpRequest, err := http.NewRequestWithContext(ctx, "GET", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("err while creating a request -> ", err)
		return
	}

	client := http.DefaultClient

	resp, err := client.Do(httpRequest)

	if err != nil {
		fmt.Println("error while making request ", err)
		return
	}

	fmt.Println("reponse -> ", resp.StatusCode)

	defer resp.Body.Close()
}
