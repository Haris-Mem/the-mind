package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Mind game!")
	})

	fmt.Println("Starting game server at :8080")
	http.ListenAndServe(":8080", nil)
}
