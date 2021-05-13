package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"l-hash-frontend/models"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ttacon/chalk"
)

func SaltGenerator() string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Register() {
	var details = models.RegisterDTO{}
	fmt.Printf("%sUsername :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Username)
	fmt.Printf("%sPassword :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%s", &details.Hash)
	fmt.Printf("%sn :: %s", chalk.Blue, chalk.Reset)
	fmt.Scanf("%d", &details.N)

	var choice rune
	fmt.Printf("\nDo you want to protect password with a salt?(Y/n) ")
	fmt.Scanf("%c", &choice)
	fmt.Println()

	if choice == 'n' || choice == 'N' {
		details.Salt = ""
	} else {
		details.Salt = SaltGenerator()
		fmt.Printf("%sSalt :: %s%s\n", chalk.Blue, chalk.Reset, details.Salt)
	}

	details.Hash = Hash(details.Hash+details.Salt, details.N)
	data, err := json.Marshal(details)

	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err.Error)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Printf("%sSuccess :: %s%s\n", chalk.Green, chalk.Reset, body)
	} else {
		fmt.Printf("%sError :: %s%s\n", chalk.Red, chalk.Reset, body)
	}
}

func Hash(hash string, n int) string {
	fmt.Printf("Hashing %s%s%s %d times...\n", chalk.Magenta, hash, chalk.Reset, n)

	for i := 0; i < n; i++ {
		hashByte := sha256.Sum256([]byte(hash))
		hash = hex.EncodeToString(hashByte[:])
	}

	fmt.Printf("%sHash :: %s%s\n", chalk.Blue, chalk.Reset, hash)
	return hash
}
