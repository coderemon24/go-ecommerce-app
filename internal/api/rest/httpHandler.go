package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpHandler struct {
	App *fiber.App	//	get the object of main fiber app ----> app := fiber.New();
	DB *gorm.DB		//	get the object of gorm db ----> db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}