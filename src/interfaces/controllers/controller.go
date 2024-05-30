package controllers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitControllers(e *echo.Echo, db *sql.DB) {
	g := e.Group("/api/v1/")
	AuthController(g, db)
	UserController(g, db)
	BranchOfficeController(g, db)
	IpsController(g, db)
	BeneficiaryController(g, db)
	EmployeeController(g, db)
	MedicalAppointmentController(g, db)
}
