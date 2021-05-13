package main

import (
	"l-hash-backend/handlers"
	"net/http"
)

// Main function
func main() {
	// Initializes all endpoints with the appropriate functions
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/login/n", handlers.GetN)
	http.HandleFunc("/reset", handlers.Reset)

	// Listen on the port 8080
	http.ListenAndServe(":8080", nil)
}
