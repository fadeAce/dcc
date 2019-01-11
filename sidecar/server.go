package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func handle(w http.ResponseWriter, req *http.Request) {
	// using DH exchanging to share secret
	if req.Method == "POST" {
		io.WriteString(w, "hello, world!\n")
	}
}

func main() {
	http.HandleFunc("/hello", handle)
	err := http.ListenAndServe(":2510", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
