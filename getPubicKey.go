package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func base64PublicKeyString() string {
	error := godotenv.Load()

	if error != nil {
		fmt.Printf("Error Loading App Key from env: %v", error)
		return "Error Loading from env"
	}

	return os.Getenv("PUBLIC_KEY")

}
