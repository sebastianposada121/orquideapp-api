package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	beneficiary_handler "orquideapp/src/interfaces/handler/beneficiary"
	beneficiary_repository "orquideapp/src/interfaces/repository/beneficiary"
	medicalAppointment_repository "orquideapp/src/interfaces/repository/medical-appointment"
	beneficiary_usecase "orquideapp/src/usecase/beneficiary"

	"github.com/labstack/echo/v4"
)

func BeneficiaryController(g *echo.Group, db *sql.DB) {
	repository := &beneficiary_repository.Repository{DB: db}
	appointmentStepRepository := &medicalAppointment_repository.Repository{DB: db}
	useCase := &beneficiary_usecase.UseCase{
		Repository:                   repository,
		MedicalAppointmentRepository: appointmentStepRepository,
	}
	handler := &beneficiary_handler.Handler{UseCase: useCase}

	bg := g.Group("beneficiaries")
	{
		mg := bg.Group("/medical-appointments")
		{
			mg.POST("", handler.CreateMedicalAppointment, middleware.JwtBeneficiaryMiddleware(db))
			mg.GET("", handler.GetAllMedicalAppointmentById, middleware.JwtBeneficiaryMiddleware(db))
		}
		cg := bg.Group("", middleware.JwtUserMiddleware(db))
		{
			cg.POST("", handler.Create)
			cg.GET("", handler.GetAll)
		}

	}

}
