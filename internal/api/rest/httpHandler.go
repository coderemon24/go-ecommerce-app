package rest

import "github.com/gofiber/fiber/v2"

type HttpHandler struct {
	App *fiber.App	//	get the object of main fiber app ----> app := fiber.New();
}