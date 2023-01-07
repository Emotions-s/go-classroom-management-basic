package main

import (
	"log"

	"github.com/Emotions-s/go-classroom-management-basic/pkg/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.Connect()
	config.GetDB()
}
