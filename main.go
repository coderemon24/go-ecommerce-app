package main

import (
	"log"

	"github.com/coderemon24/go-ecommerce-app/config"
	"github.com/coderemon24/go-ecommerce-app/internal/api"
)


func main() {
	cfg, err := config.SetupEnv()
	
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	
	api.NewServer(cfg)
	
}