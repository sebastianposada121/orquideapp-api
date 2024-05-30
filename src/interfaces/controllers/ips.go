package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	ips_handler "orquideapp/src/interfaces/handler/ips"
	ips_repository "orquideapp/src/interfaces/repository/ips"
	ips_usecase "orquideapp/src/usecase/ips"

	"github.com/labstack/echo/v4"
)

func IpsController(g *echo.Group, db *sql.DB) {
	repository := &ips_repository.Repository{DB: db}
	useCase := &ips_usecase.UseCase{Repository: repository}
	handler := &ips_handler.Handler{UseCase: *useCase}

	ig := g.Group("ips", middleware.JwtUserMiddleware(db))
	{
		ig.POST("", handler.Create)
		ig.GET("", handler.GetAllByBranchOfficeID)
	}

}
