package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"l-hash-backend/database"
	"l-hash-backend/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/ttacon/chalk"
)

// Validate is used to check if hash(given_hash) = hash in database
func Validate(username, hash string, userCreds database.UserCreds) (bool, string) {
	// Take SHA256 hash of the hash given by user
	hashOfHash := sha256.Sum256([]byte(hash))
	hashString := hex.EncodeToString(hashOfHash[:])
	// If the hash is not equal throw error
	if hashString != userCreds.Hash {
		fmt.Printf("%s Error ::%s Hash is Invalid\n\n", chalk.Red, chalk.Reset)
		return false, "Invalid Hash. Your password was incorrect."
	}

	// If equal decrement n and change hash in database
	userCreds.Hash = hash
	userCreds.N -= 1
	database.Set(username, userCreds)
	return true, ""
}

// Login checks if given hash of user is valid
func Login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Login --"), chalk.Reset)
	var p models.LoginDTO

	// Decodes user input to JSON
	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	// Validates if all fields are present in the POST request
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error ::%s Invalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n", p.Username, p.Hash)
	// Check if user is present in the database
	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	// Fetch the user details from the database
	userCreds := database.Get(p.Username)
	// If N value is 2, prompt user to reset his password immediately
	if userCreds.N == 2 {
		fmt.Printf("%s Error ::%s Reset Password Immediately\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Reset Password Immediately", http.StatusBadRequest)
		return
	}

	// Check if the hash given by user is valid
	ok, errString := Validate(p.Username, p.Hash, userCreds)
	if !ok {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	// If valid return success message
	fmt.Printf("%s Success\n\n%s", chalk.Green, chalk.Reset)
	fmt.Fprintf(w, "User Successfully logged in.")
}

// GetN returns the N and salt of a given username
func GetN(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Get N --"), chalk.Reset)

	var p models.GetNDTO
	// Decodes user input to JSON
	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	// Validates if all fields are present in the POST request
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error :: %sInvalid Data", chalk.Red, chalk.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n", p.Username)

	// Check if user is present in the database
	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	// Fetch the user details from the database
	details := database.Get(p.Username)
	data := models.GetNResponse{N: details.N, Salt: details.Salt}

	// Return the user details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
}
