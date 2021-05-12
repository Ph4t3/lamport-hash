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
	if err != nil || p.Hash == "" || p.Username == "" || p.N == 0 {
		fmt.Printf("%s Error :: %sInvalid Data\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	fmt.Printf(" Username :: %s \n Hash :: %s\n n :: %d\n", p.Username, p.Hash, p.N)

	if database.Check(p.Username) {
		fmt.Printf("%s Error :: %sUser already registered\n\n", chalk.Red, chalk.Reset)
		http.Error(w, "User already Registered", http.StatusBadRequest)
		return
	}

	userCreds := database.UserCreds{Hash: p.Hash, N: p.N}
	database.Set(p.Username, userCreds)

	fmt.Printf("%s Success\n\n%s", chalk.Blue, chalk.Reset)
	fmt.Fprintf(w, "User successfully registered")
}
