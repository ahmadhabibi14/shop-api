package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func main() {
	id := generateRandomID()
	fmt.Println(id)
}

func generateRandomID() string {
	// Generate a 16-byte random number
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(err)
	}

	// Encode the random number to a base64 string
	encode := base64.StdEncoding.EncodeToString(b)
	replacer := strings.NewReplacer(
		"&", "",
		"-", "",
		"+", "",
		"=", "",
		"!", "",
		"/", "",
		`\`, "",
		"#", "",
		"*", "",
		"%", "",
	)
	id := replacer.Replace(encode)
	return id
}
