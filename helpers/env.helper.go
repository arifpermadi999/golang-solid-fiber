package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
}
