package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type timeout interface {
	Timeout() bool
}

func main() {
	client := http.Client{
		Timeout: 4 * time.Second,
	}
	//res, err := client.Get("http://localhost:8081/hello")

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/hello", nil)
	fmt.Println(err)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(timeout); ok {
			fmt.Println("timeout", err.Timeout())
		}
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(data))

}
