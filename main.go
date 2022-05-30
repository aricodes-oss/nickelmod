package main

import (
	"nickelmod/logging"

	"github.com/joho/godotenv"
)

func init() {
	// Disregarding error from this because we might not have a .env file
	godotenv.Load()
}

func main() {
	log := logging.Default

	log.Info("Hello, world!")
}
