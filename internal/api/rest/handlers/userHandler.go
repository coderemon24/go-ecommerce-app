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
	app.Post("/users/login", handler.UserLogin)
	app.Post("/users/register", handler.UserRegister)
	
	
	// private endpoints
	app.Get("/users", handler.UserList)
	app.Post("/users", handler.UserCreate)
	app.Put("/users/:id", handler.UserUpdate)
	app.Delete("/users/:id", handler.UserDelete)
	app.Get("/users/:id", handler.UserById)
	app.Get("/users/email/:email", handler.UserByEmail)
}


// receiver functions -> methods

// user list
func (h *UserHandler) UserList(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "user list",
	})
}
// user by id
func (h *UserHandler) UserById(ctx *fiber.Ctx) error{
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user by id",
	})
}
// user by email
func (h *UserHandler) UserByEmail(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "user by email",
	})
}
// user create
func (h *UserHandler) UserCreate(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user created",
	})
}
// user update
func(h *UserHandler)UserUpdate(ctx *fiber.Ctx)error{
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "user updated",
	})
}
// user delete
func(h *UserHandler) UserDelete(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message" : "user deleted",
	})
}
// user login
func (h *UserHandler) UserLogin(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "user logged in",
	})
}
// user register
func (h *UserHandler) UserRegister(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "user registered",
	})
}