package routes

import (
	"backend/controller"
	"backend/repository"
	"backend/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupEmployeeRoutes(db *gorm.DB, router fiber.Router) {
	employee := router.Group("/employee")

	er := repository.NewEmployeeRepository(db)
	eu := usecase.NewEmployeeUsecase(er)
	ec := controller.NewEmployeeController(eu)

	employee.Get("/", ec.GetAllEmployees)
	employee.Get("/:employeeId", ec.GetEmployeeById)
}
