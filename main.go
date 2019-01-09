package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw, `{"message": "Hello World"}`)
	})

	fmt.Println("Starting simple http server...")

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
