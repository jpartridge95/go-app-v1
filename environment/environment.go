package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading env File")
	}
	return os.Getenv(key)
}
