package routes


import (
	"github.com/gofiber/fiber/v2"
	"app_food/internal/httpserver/controllers"
	"app_food/pkg/middlewares"

)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/auth")

	route.Post("/signin", controllers.LoginController)         
	route.Post("/signup", controllers.SignUpController)         
	route.Post("/refresh", middlewares.JwtRefresh(), controllers.RefreshController)         

}