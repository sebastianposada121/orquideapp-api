package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	beneficiary_handler "orquideapp/src/interfaces/handler/beneficiary"
	appointmentstep_repository "orquideapp/src/interfaces/repository/appointment-step"
	beneficiary_repository "orquideapp/src/interfaces/repository/beneficiary"
	beneficiary_usecase "orquideapp/src/usecase/beneficiary"

	"github.com/labstack/echo/v4"
)

func BeneficiaryController(g *echo.Group, db *sql.DB) {
	repository := &beneficiary_repository.Repository{DB: db}
	appointmentStepRepository := &appointmentstep_repository.Repository{DB: db}

	useCase := &beneficiary_usecase.UseCase{
		Repository:                repository,
		AppointmentStepRepository: appointmentStepRepository,
	}

	handler := &beneficiary_handler.Handler{UseCase: useCase}

	bg := g.Group("beneficiaries")
	{
		bg.POST("/create-medical-appointment", handler.CreateMedicalAppointment, middleware.JwtBeneficiaryMiddleware(db))
		cg := bg.Group("", middleware.JwtUserMiddleware(db))
		{
			cg.POST("", handler.Create)
			cg.GET("", handler.GetAll)
		}

	}

}
