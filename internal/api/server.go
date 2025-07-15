package api

import (
	"github.com/coderemon24/go-ecommerce-app/config"
	"github.com/coderemon24/go-ecommerce-app/internal/api/rest"
	"github.com/coderemon24/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func NewServer(cfg config.AppConfig) {
	server := cfg.Host + ":" + cfg.Port
	
	app := fiber.New();
	
	/**********************************************************************
	* => instance of HttpHandler
	* => passing fiber app through the HttpHandler instance of struct
	***********************************************************************/
	
	rh := &rest.HttpHandler{
		App: app,
	}
	
	setupRoutes(rh)
	
	
	app.Listen(server)
}

//	registering routes via HttpHandler struct using receiver function
func setupRoutes(rh *rest.HttpHandler) {
	handlers.UserRoutes(rh)
}