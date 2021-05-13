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

func Validate(username, hash string, userCreds database.UserCreds) (bool, string) {
	hashOfHash := sha256.Sum256([]byte(hash))
	hashString := hex.EncodeToString(hashOfHash[:])
	if hashString != userCreds.Hash {
		fmt.Printf("%s Error ::%s Hash is Invalid\n\n", chalk.Red, chalk.Reset)
		return false, "Invalid Hash. Your password was incorrect."
	}

	userCreds.Hash = hash
	userCreds.N -= 1
	database.Set(username, userCreds)
	return true, ""
}

func Login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Login --"), chalk.Reset)
	var p models.LoginDTO

	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error ::%s Invalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n", p.Username, p.Hash)
	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	userCreds := database.Get(p.Username)
	if userCreds.N == 2 {
		fmt.Printf("%s Error ::%s Reset Password Immediately\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Reset Password Immediately", http.StatusBadRequest)
		return
	}

	ok, errString := Validate(p.Username, p.Hash, userCreds)
	if !ok {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	fmt.Printf("%s Success\n\n%s", chalk.Green, chalk.Reset)
	fmt.Fprintf(w, "User Successfully logged in.")
}

func GetN(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Get N --"), chalk.Reset)

	var p models.GetNDTO
	err := json.NewDecoder(req.Body).Decode(&p)
	validate := validator.New()
	if err != nil || validate.Struct(p) != nil {
		fmt.Printf("%s Error :: %sInvalid Data", chalk.Red, chalk.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n", p.Username)

	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	details := database.Get(p.Username)
	data := models.GetNResponse{details.N, details.Salt}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
}
