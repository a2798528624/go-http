package main

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello") // 服务端打印输出
	time.Sleep(2 * time.Second)
	fmt.Println("1213")
	fmt.Fprintf(w, "hello GoLangWEB")
}

func main() {
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	fmt.Println(err)

}
