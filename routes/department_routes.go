package routes

import (
	"backend/controller"
	"backend/repository"
	"backend/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupDepartmentRoutes(db *gorm.DB, router fiber.Router) {
	department := router.Group("/department")

	dr := repository.NewDepartmentRepository(db)
	du := usecase.NewDepartmentUsecase(dr)
	dc := controller.NewDepartmentController(du)

	department.Get("/", dc.GetAllDepartments)
	department.Get("/:departmentId", dc.GetDepartmentById)
}
