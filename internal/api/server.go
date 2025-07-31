package api

import (
	"log"

	"github.com/coderemon24/go-ecommerce-app/config"
	"github.com/coderemon24/go-ecommerce-app/internal/api/rest"
	"github.com/coderemon24/go-ecommerce-app/internal/api/rest/handlers"
	"github.com/coderemon24/go-ecommerce-app/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewServer(cfg config.AppConfig) {
	server := cfg.Host + ":" + cfg.Port
	dsn := cfg.DSN
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	log.Println("Database connected")
	
	db.AutoMigrate(&domain.User{})
	
	
	app := fiber.New();
	
	/**********************************************************************
	* => instance of HttpHandler
	* => passing fiber app through the HttpHandler instance of struct
	***********************************************************************/
	
	rh := &rest.HttpHandler{
		App: app,
		DB:  db,
	}
	
	setupRoutes(rh)
	
	
	app.Listen(server)
}

//	registering routes via HttpHandler struct using receiver function
func setupRoutes(rh *rest.HttpHandler) {
	handlers.UserRoutes(rh)
	handlers.CatalogRoutes(rh)
}