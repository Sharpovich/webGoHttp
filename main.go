package main

import (
	"log"
	"os"
	"project/com"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	com.Routers(os.Getenv("WEB_HOST"), os.Getenv("WEB_PORT"))
}
