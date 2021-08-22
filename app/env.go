package app

import (
	"log"

	"github.com/joho/godotenv"
)

// Load all environmental variables from .env into the current environment.
func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No env file found.")
	}
}
