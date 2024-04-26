package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func LoadEnv() {
	err := godotenv.Load("/home/vladyslav/Documents/web-projects/greenlight/.env")
	if err != nil {
		fmt.Println("Eror loading .env")
	}

}
func GetDBConfig() string {
	LoadEnv()
	connStr := os.Getenv("GREENLIGHT_DB_DSN")
	return connStr
}
