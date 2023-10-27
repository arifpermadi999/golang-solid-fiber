package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(database *gorm.DB) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Hello,World")
	})
	AuthRoute(database, v1)
	CategoryRoute(database, v1)

	return app
}
