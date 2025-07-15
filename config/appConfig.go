package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host string
	Port string
}

func SetupEnv() (cfg AppConfig, err error) {

	err = godotenv.Load()
	
	if  err != nil {
		return AppConfig{}, errors.New("error loading .env file")
	}
	
	
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	
	if port == "" {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}
	if host == "" {
		return AppConfig{}, errors.New("HTTP_HOST is not set")
	}
	
	return AppConfig{
		Host: host,
		Port: port,
	}, nil
	
}