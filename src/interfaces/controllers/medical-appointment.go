package controllers

import (
	"database/sql"
	medicalappointment_handler "orquideapp/src/interfaces/handler/medical-appointment"
	medicalAppointment_repository "orquideapp/src/interfaces/repository/medical-appointment"
	medicalAppointment_usecase "orquideapp/src/usecase/medical-appointment"

	"github.com/labstack/echo/v4"
)

func MedicalAppointmentController(g *echo.Group, db *sql.DB) {
	repository := medicalAppointment_repository.Repository{DB: db}
	usecase := medicalAppointment_usecase.UseCase{Repository: &repository}
	handler := medicalappointment_handler.Handler{UseCase: &usecase}

	mg := g.Group("medical-appointments")
	{
		mg.GET("/steps/:id", handler.GetAllAppointmentStepsByMedicalAppointmentID)
	}
}
