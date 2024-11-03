package main

import (
	"app_food/internal/httpserver/routes"
	"app_food/pkg/config"
	"app_food/pkg/log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Backend Web server
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	
	app.Use(swagger.New(cfg))
	routes.AuthRoutes(app)
	routes.NotFoundRoute(app)

	port, _ := config.GetConfig("PORT")
	if port == "" {
		port = "3000" 
		log.Logger.Warn("No PORT specified, defaulting to 3000")
	}

	if err := app.Listen(":" + port); err != nil {
		log.Logger.Errorf("Error: app failed to start on port %s, %v", port, err)
		panic(err)
	}

}

