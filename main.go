package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from GitLab CI/CD!!")
}

// Handler function for testing endpoint
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from test endpoint"))
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/test", TestHandler)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
