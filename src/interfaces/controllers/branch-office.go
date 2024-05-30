package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	branchOffice_hander "orquideapp/src/interfaces/handler/branch-office"
	branchOffice_repository "orquideapp/src/interfaces/repository/branch-office"
	branchOffice_usecase "orquideapp/src/usecase/branch-office"

	"github.com/labstack/echo/v4"
)

func BranchOfficeController(g *echo.Group, db *sql.DB) {
	repository := &branchOffice_repository.Repository{DB: db}
	useCase := &branchOffice_usecase.UseCase{Repository: repository}
	handler := &branchOffice_hander.Handler{UseCase: *useCase}

	eg := g.Group("branch-offices", middleware.JwtUserMiddleware(db))
	{
		eg.POST("", handler.Create)
		eg.GET("", handler.GetAllByUserId)
	}

}
