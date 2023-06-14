package routes

import (
	"backend/bootstrap"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *bootstrap.App, fiberApp *fiber.App) {
	api := fiberApp.Group("/api")
	v1 := api.Group("/v1")

	setupEmployeeRoutes(app.DB, v1)
	setupDepartmentRoutes(app.DB, v1)
}
