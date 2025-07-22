package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error when loading file configuration: %v", err.Error())
	}

	expInt, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		return &Config{}, errors.New("failed to convert integer to string")
	}

	return &Config{
		Server: Server{
			Host:  os.Getenv("SERVER_HOST"),
			Port:  os.Getenv("SERVER_PORT"),
			Asset: os.Getenv("SERVER_ASSET_URL"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Jwt: Jwt{
			Key: os.Getenv("JWT_KEY"),
			Exp: expInt,
		},
		Storage: Storage{
			BasePath: os.Getenv("STORAGE_PATH"),
		},
	}, nil
}
