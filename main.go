package main

import (
	"github.com/joho/godotenv"
	"log"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error while loading .env file: " + err.Error())
	}
}

func main() {
	loadenv()
}
