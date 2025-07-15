package handlers

import (
	"net/http"

	"github.com/coderemon24/go-ecommerce-app/internal/api/rest"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// service user.Service
}

func UserRoutes(rh *rest.HttpHandler) {
	app := rh.App
	
	// instance of UserHandler Struct
	handler := UserHandler{}
	
	// public endpoints
	app.Get("/users", handler.UserList)
	
	// private endpoints
}


// receiver function
func (h *UserHandler) UserList(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "user list",
	})
}
