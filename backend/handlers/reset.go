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

// Reset is used to reset the password of a user
func Reset(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Password Reset --"), chalk.Reset)
	var p models.ResetDTO

	// Decodes user input to JSON
	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	// Validates if all fields are present in the POST request
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error ::%s Invalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(
		" Username :: %s \n Hash :: %s\n New Hash :: %s\n Salt :: %s\n n :: %d\n",
		p.Username,
		p.Hash,
		p.Newhash,
		p.Salt,
		p.N,
	)

	// Check if user is present in the database
	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	// Fetch the user details from the database
	userCreds := database.Get(p.Username)
	// Check if the hash given by the user is valid
	ok, errString := Validate(p.Username, p.Hash, userCreds)
	if !ok {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	// Update the new hash and N in the database
	userCreds.Hash = p.Newhash
	userCreds.N = p.N
	userCreds.Salt = p.Salt
	database.Set(p.Username, userCreds)

	// Return success message
	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
	fmt.Fprintf(w, "Password Reset Successfully")
}
