package controller

import (
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController interface {
	GetAllEmployees(c *fiber.Ctx) error
	GetEmployeeById(c *fiber.Ctx) error
}

type employeeController struct {
	eu usecase.EmployeeUsecase
}

func NewEmployeeController(tu usecase.EmployeeUsecase) EmployeeController {
	return &employeeController{tu}
}

func (ec *employeeController) GetAllEmployees(c *fiber.Ctx) error {
	employeesRes, err := ec.eu.GetAllEmployees()

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(employeesRes)
}

func (ec *employeeController) GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("employeeId")
	employeeId, _ := strconv.Atoi(id)
	employeeRes, err := ec.eu.GetEmployeeById(uint(employeeId))

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(employeeRes)
}
