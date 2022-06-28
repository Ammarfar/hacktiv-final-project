package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(param string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(param)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
