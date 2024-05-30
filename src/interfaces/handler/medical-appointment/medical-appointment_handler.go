package medicalappointment_handler

import (
	"net/http"
	medicalAppointment_usecase "orquideapp/src/usecase/medical-appointment"
	"orquideapp/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UseCase *medicalAppointment_usecase.UseCase
}

func (h *Handler) GetAllAppointmentStepsByMedicalAppointmentID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "no found id", nil)
	}

	appointmentSteps, err := h.UseCase.GetAllAppointmentStepsByMedicalAppointmentID(id)

	if err != nil {
		return utils.GenericResponse(c, http.StatusBadRequest, false, "error get steps", nil)
	}

	return utils.GenericResponse(c, http.StatusAccepted, true, "Success", appointmentSteps)
}
