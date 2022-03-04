package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":50051", nil))
}
