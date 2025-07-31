package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host string
	Port string
	DSN  string
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
	
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_ssl_mode := os.Getenv("DB_SSL_MODE")
	
	if db_host == "" {
		return AppConfig{}, errors.New("DB_HOST is not set")
	}
	if db_port == "" {
		return AppConfig{}, errors.New("DB_PORT is not set")
	}
	if db_name == "" {
		return AppConfig{}, errors.New("DB_NAME is not set")
	}
	if db_username == "" {
		return AppConfig{}, errors.New("DB_USERNAME is not set")
	}
	if db_password == "" {
		return AppConfig{}, errors.New("DB_PASSWORD is not set")
	}
	dsn := "host="+db_host+" port="+db_port+" user="+db_username+" password="+db_password+" dbname="+db_name+" sslmode="+db_ssl_mode
	
	
	return AppConfig{
		Host: host,
		Port: port,
		DSN:  dsn,
	}, nil
	
}