package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	employee_handler "orquideapp/src/interfaces/handler/employee"
	employee_repository "orquideapp/src/interfaces/repository/employee"
	employee_usecase "orquideapp/src/usecase/employee"

	"github.com/labstack/echo/v4"
)

func EmployeeController(g *echo.Group, db *sql.DB) {
	repository := &employee_repository.Repository{DB: db}
	useCase := &employee_usecase.UseCase{Repository: repository}
	handler := &employee_handler.Handler{UseCase: *useCase}

	eg := g.Group("employees")
	{
		eg.POST("/roles", handler.CreateRol, middleware.JwtEmployeMiddleware(db))
		eg.POST("", handler.Create, middleware.JwtUserMiddleware(db))
	}
}