package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)

	fmt.Println("JWT Key generated:")
	fmt.Println(encodedKey)
}
