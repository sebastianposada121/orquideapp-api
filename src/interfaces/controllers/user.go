package controllers

import (
	"database/sql"
	"orquideapp/middleware"
	user_handler "orquideapp/src/interfaces/handler/user"
	user_repository "orquideapp/src/interfaces/repository/user"
	user_usecase "orquideapp/src/usecase/user"

	"github.com/labstack/echo/v4"
)

func UserController(g *echo.Group, db *sql.DB) {
	repository := &user_repository.Repository{DB: db}
	useCase := &user_usecase.UseCase{Repository: repository}
	handler := &user_handler.Handler{UseCase: useCase}
	ug := g.Group("users")
	{
		ug.GET("/:email", handler.GetByEmail, middleware.JwtUserMiddleware(db))
	}

}
