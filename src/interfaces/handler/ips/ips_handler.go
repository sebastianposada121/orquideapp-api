package ips_handler

import (
	"net/http"
	"orquideapp/src/domain"
	ips_usecase "orquideapp/src/usecase/ips"
	"orquideapp/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase ips_usecase.UseCase
}

func (h *Handler) Create(c echo.Context) error {
	ips := new(domain.Ips)

	if err := c.Bind(ips); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error creating ips", nil)
	}

	if err := h.UseCase.Create(*ips); err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error creating ips", nil)
	}

	return utils.GenericResponse(c, http.StatusCreated, true, "ips created successfully", ips)
}

func (h *Handler) GetAllByBranchOfficeID(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("branchOfficeId"))
	user := c.Get("user").(domain.User)
	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error id", nil)
	}

	ipsList, err := h.UseCase.GetAllByBranchOfficeID(user.ID, id)

	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error", nil)
	}

	return utils.GenericResponse(c, http.StatusAccepted, true, "get all branch offices", ipsList)
}

func (h *Handler) GetByID(c echo.Context) error {
	return utils.GenericResponse(c, http.StatusAccepted, true, "get all branch", nil)
}
