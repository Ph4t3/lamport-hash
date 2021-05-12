package handlers

import (
	"encoding/json"
	"fmt"
	database "l-hash-backend/database"
	"net/http"

	"github.com/ttacon/chalk"
)

func Register(w http.ResponseWriter, req *http.Request) {
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("-- Register --"), chalk.Reset)

	var p RegisterDTO
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil || p.Hash == "" || p.Username == "" {
		fmt.Printf("%s Error :: %sInvalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n", p.Username, p.Hash)

	if database.Check(p.Username) {
		fmt.Printf("%s Error :: %sUser already registered\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "User already Registered", http.StatusBadRequest)
		return
	}

	userCreds := database.UserCreds{Hash: p.Hash, N: 100}
	database.Set(p.Username, userCreds)

	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
	fmt.Fprintf(w, "Success")
}
