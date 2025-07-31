package handlers

import (
	"net/http"

	"github.com/coderemon24/go-ecommerce-app/internal/api/rest"
	"github.com/coderemon24/go-ecommerce-app/internal/dto"
	"github.com/coderemon24/go-ecommerce-app/internal/repository"
	"github.com/coderemon24/go-ecommerce-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func UserRoutes(rh *rest.HttpHandler) {
	app := rh.App
	
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	
	// instance of UserHandler Struct
	handler := UserHandler{
		svc: svc,
	}
	
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
	
	email := ctx.Params("email")
	
	user, err := h.svc.FindUserByEmail(email)
	if err != nil {
		return ctx.Status(404).JSON(&fiber.Map{
			"message": "user not found",
		})
	}
	
	return ctx.Status(200).JSON(&fiber.Map{
		"email": user,
		"message": "user by email",
	})
}
// user create
func (h *UserHandler) UserCreate(ctx *fiber.Ctx) error {
	
	usr := dto.UserRegister{}
	err := ctx.BodyParser(&usr)
	
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request",
		})
	}
	
	_, err = h.svc.Signup(usr)
	
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "failed to create user",
		})
	}
	
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"user": usr,
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
	
	user := dto.UserRegister{}
	err := ctx.BodyParser(&user)
	
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "invalid request",
		})
	}
	
	_, err = h.svc.Signup(user)
	
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"message": "user registration failed",
		})
	}
	
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "user registered",
	})
}