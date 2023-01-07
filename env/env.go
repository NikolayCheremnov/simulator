package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
}

func GetPortStr() string {
	return os.Getenv("APP_PORT")
}

func GetPort() int {
	port, err := strconv.Atoi(GetPortStr())
	if err != nil {
		panic("error loading .env file: NO APP_PORT")
	}
	return port
}
