package main

import (
	"github.com/joho/godotenv"
	"log"
	"usermanagement/usecase"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ..env file not FOUND")
	}

	usecase.Init()
}
