package routes

import (
	"github.com/arifpermadi999/core-echo-golang/injection"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(db *gorm.DB, routes fiber.Router) {
	AuthApi := injection.InitAuthApi(db)
	//r := routes.Group("")
	routes.Post("/login", AuthApi.Login)
	routes.Post("/register", AuthApi.Register)
}
func CategoryRoute(db *gorm.DB, routes fiber.Router) {
	CategoryApi := injection.InitCategoryApi(db)
	r := routes.Group("category")
	r.Get("/", CategoryApi.GetAllCategory)
	r.Put("/:id", CategoryApi.ViewCategory)
	r.Delete("/:id", CategoryApi.DeleteCategory)
	r.Post("/update", CategoryApi.UpdateCategory)
	r.Post("/save", CategoryApi.SaveCategory)
}
