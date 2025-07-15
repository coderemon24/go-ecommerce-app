package api

import (
	"github.com/coderemon24/go-ecommerce-app/config"
	"github.com/gofiber/fiber/v2"
)

func NewServer(cfg config.AppConfig) {
	server := cfg.Host + ":" + cfg.Port
	
	app := fiber.New();
	
	app.Listen(server)
}