package controller

import (
	"backend/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DepartmentController interface {
	GetAllDepartments(c *fiber.Ctx) error
}

type departmentController struct {
	du usecase.DepartmentUsecase
}

func NewDepartmentController(du usecase.DepartmentUsecase) DepartmentController {
	return &departmentController{du}
}

func (dc *departmentController) GetAllDepartments(c *fiber.Ctx) error {
	departmentsRes, err := dc.du.GetAllDepartments()

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(departmentsRes)
}
