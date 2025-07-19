package handlers

import (
	"net/http"

	"github.com/coderemon24/go-ecommerce-app/internal/api/rest"
	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	// service catalog.Service
}

func CatalogRoutes(rh *rest.HttpHandler){
	app := rh.App
	
	handlers := CatalogHandler{}
	
	// public endpoints
	app.Get("/catalog", handlers.CatalogList)
	
	// private endpoints
	app.Post("/catalog", handlers.CatalogCreate)
}


func (h *CatalogHandler) CatalogList(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "catalog list",
	})
}

func (h *CatalogHandler) CatalogCreate(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "catalog created",
	})
}