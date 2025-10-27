package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func app_key() string {

	error := godotenv.Load()

	if error != nil {
		fmt.Printf("Error Loading App Key from env: %v", error)
		return "Error Loading from env"
	}

	return os.Getenv("APP_KEY")
}
