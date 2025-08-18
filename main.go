package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Welcome to the Go HTTP server.")
}

func envEchoHandler(w http.ResponseWriter, r *http.Request) {
	// Get an environment variable (using USER as an example, but you can change this)
	user := os.Getenv("USER")
	if user == "" {
		user = "Unknown"
	}

	message := fmt.Sprintf("Hello %s! This message contains an environment variable: USER=%s", user, user)
	fmt.Fprintf(w, message)
}

func secretEchoHandler(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("URL")
	if secret == "" {
		secret = "Unknown"
	}
	message := fmt.Sprintf("Hello %s! This message contains a secret: URL=%s", secret, secret)
	fmt.Fprintf(w, message)
}
func main() {
	// Define routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/env-echo", envEchoHandler)
	http.HandleFunc("/secret-echo", secretEchoHandler)

	// Set port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s...\n", port)
	fmt.Printf("Available endpoints:\n")
	fmt.Printf("  - GET / or /hello: Say hello\n")
	fmt.Printf("  - GET /env-echo: Echo message with environment variable\n")

	// Start the server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
