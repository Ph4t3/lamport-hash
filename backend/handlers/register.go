package handlers

import (
	"encoding/json"
	"fmt"
	"l-hash-backend/database"
	"l-hash-backend/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/ttacon/chalk"
)

// Register is used to register a new user in the database
func Register(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Register --"), chalk.Reset)

	var p models.RegisterDTO
	// Decodes user input to JSON
	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	// Validates if all fields are present in the POST request
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error :: %sInvalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n Salt :: %s\n n :: %d\n", p.Username, p.Hash, p.Salt, p.N)

	// Check if user is present in the database
	if database.Check(p.Username) {
		fmt.Printf("%s Error :: %sUser already registered\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "User already Registered", http.StatusBadRequest)
		return
	}

	// Create new user and add to the database
	userCreds := database.UserCreds{Hash: p.Hash, N: p.N, Salt: p.Salt}
	database.Set(p.Username, userCreds)

	// Return success message
	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
	fmt.Fprintf(w, "User successfully registered")
}
