package controllers

import (
	"database/sql"
	auth_handler "orquideapp/src/interfaces/handler/auth"
	beneficiary_repository "orquideapp/src/interfaces/repository/beneficiary"
	employee_repository "orquideapp/src/interfaces/repository/employee"
	user_repository "orquideapp/src/interfaces/repository/user"
	auth_usecase "orquideapp/src/usecase/auth"

	"github.com/labstack/echo/v4"
)

func AuthController(g *echo.Group, db *sql.DB) {
	useCase := &auth_usecase.UseCase{
		UserRepository:        &user_repository.Repository{DB: db},
		EmployeeRepository:    &employee_repository.Repository{DB: db},
		BeneficiaryRepository: &beneficiary_repository.Repository{DB: db},
	}
	authHandler := &auth_handler.Handler{UseCase: useCase}

	users := g.Group("auth")
	{
		users.POST("/login", authHandler.UserLogin)
		users.POST("/signup", authHandler.UserSignup)
		users.POST("/login-e", authHandler.EmployeeLogin)
	}

	employees := g.Group("employees")
	{
		auth := employees.Group("/auth")
		{
			auth.POST("/login", authHandler.EmployeeLogin)
			auth.PUT("/update-password", authHandler.UpdateEmployeePassword)
		}
	}

	benefeciary_g := g.Group("beneficiaries")
	{
		auth := benefeciary_g.Group("/auth")
		{
			auth.POST("/login", authHandler.BeneficiaryLogin)
			auth.PUT("/update-password", authHandler.UpdateBeneficiaryPassword)
		}
	}
}
