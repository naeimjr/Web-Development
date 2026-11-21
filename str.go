package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Habib. I'm youtuber. I'm software engineer.")
}

func main() {
	mux := http.NewServeMux() // Router

	mux.HandleFunc("/hello", helloHandler) // Route
	mux.HandleFunc("/about", aboutHandler) // Route

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
