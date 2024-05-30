package employee_handler

import (
	"net/http"
	"orquideapp/src/domain"
	employee_usecase "orquideapp/src/usecase/employee"
	"orquideapp/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase employee_usecase.UseCase
}

func (h *Handler) Create(c echo.Context) error {
	employee := new(domain.Employee)
	if err := c.Bind(employee); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create employee", nil)
	}
	if err := h.UseCase.Create(*employee); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error create employee", nil)
	}
	return utils.GenericResponse(c, http.StatusCreated, true, "employee create successfully", nil)
}

func (h *Handler) CreateRol(c echo.Context) error {
	rol := new(domain.Rol)
	employee := c.Get("employee").(domain.Employee)

	if err := c.Bind(rol); err != nil {
		return utils.GenericResponse(c, http.StatusCreated, true, "error create", nil)
	}
	rol.IpsId = employee.IpsId
	return utils.GenericResponse(c, http.StatusCreated, true, "rol create successfully", nil)
}
