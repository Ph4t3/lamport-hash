package handlers

import (
	"encoding/json"
	"fmt"
	database "l-hash-backend/database"
	"net/http"

	"github.com/ttacon/chalk"
)

func Reset(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Password Reset --"), chalk.Reset)
	var p ResetDTO

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil || p.Username == "" || p.Hash == "" || p.Newhash == "" || p.N == 0 {
		fmt.Printf("%s Error ::%s Invalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n New Hash :: %s\n n :: %d\n", p.Username, p.Hash, p.Newhash, p.N)
	if !database.Check(p.Username) {
		fmt.Printf("%s Error ::%s User does not exist\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Username", http.StatusBadRequest)
		return
	}

	userCreds := database.Get(p.Username)
	ok, errString := Validate(p.Username, p.Hash, userCreds)
	if !ok {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	userCreds.Hash = p.Newhash
	userCreds.N = p.N
	database.Set(p.Username, userCreds)

	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
	fmt.Fprintf(w, "success")
}
