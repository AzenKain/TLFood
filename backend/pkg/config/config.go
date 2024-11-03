package config

import (
	"errors"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GetConfig(config string) (string, error) {
	var data  string =  os.Getenv(config)
	if (data == "") {
		return "", errors.New("this config dose not exit")
	}

	return data, nil
} 


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}
}