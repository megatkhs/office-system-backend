package controller

import (
	"backend/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController interface {
	GetAllEmployees(c *fiber.Ctx) error
}

type employeeController struct {
	eu usecase.EmployeeUsecase
}

func NewEmployeeController(tu usecase.EmployeeUsecase) EmployeeController {
	return &employeeController{tu}
}

func (ec *employeeController) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := ec.eu.GetAllEmployees()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(employees)
}
