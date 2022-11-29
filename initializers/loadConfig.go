package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error while loading env loading .env file")
	}
}
