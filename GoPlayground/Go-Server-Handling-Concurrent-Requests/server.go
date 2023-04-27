package main

import (
	"fmt"
	"net/http"
	"time"
)

var count int64

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler Says Hello")
	time.Sleep(500 * time.Millisecond)
	w.WriteHeader(200)
	r.Body.Close()
	count++
}

func main() {
	http.HandleFunc("/counter", handle)
	fmt.Println("Starting Server:")
	if err := http.ListenAndServe(":7777", nil); err != nil {
		panic(err)
	}
}
