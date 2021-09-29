package main

import (
	"log"
	"os"
	"project/apps"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	apps.Routers(os.Getenv("WEB_HOST"), os.Getenv("WEB_PORT"))
}
