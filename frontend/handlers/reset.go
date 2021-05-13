package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"l-hash-frontend/models"
	"log"
	"net/http"

	"github.com/ttacon/chalk"
)

func Reset() {
	var details = models.ResetDTO{}
	fmt.Printf("%sUsername :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Username)
	fmt.Printf("%sOld Password :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Hash)
	fmt.Printf("%sNew Password :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Newhash)
	fmt.Printf("%sn :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%d", &details.N)
	fmt.Println()

	ok, n, salt := GetN(details.Username)
	if !ok {
		return
	}

	fmt.Printf("%sHashing Old Password%s\n", chalk.Magenta, chalk.Reset)
	details.Hash = Hash(details.Hash+salt, n-1)
	fmt.Printf("\n%sHashing New Password%s\n", chalk.Magenta, chalk.Reset)
	details.Salt = SaltGenerator()
	fmt.Printf("%sNew Salt :: %s%s\n", chalk.Blue, chalk.Reset, details.Salt)
	details.Newhash = Hash(details.Newhash+details.Salt, details.N)
	data, err := json.Marshal(details)

	resp, err := http.Post("http://localhost:8080/reset", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err.Error)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Printf("\n%sSuccess :: %s%s\n", chalk.Green, chalk.Reset, body)
	} else {
		fmt.Printf("%sError :: %s%s\n", chalk.Red, chalk.Reset, body)
	}
}
