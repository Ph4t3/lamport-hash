package main

import (
	"l-hash-backend/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/login/n", handlers.GetN)
	http.HandleFunc("/reset", handlers.Reset)

	//fmt.Print(" =============================================== \n")
	//fmt.Print("|						|\n")
	//fmt.Print("|    App running at: http://localhost:8080	|\n")
	//fmt.Print("|						|\n")
	//fmt.Print(" =============================================== \n\n")

	http.ListenAndServe(":8080", nil)
}
