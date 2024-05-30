package domain

type AuthRepository struct {
	userRepo        UserRepository
	employeeRepo    EmployeesRepository
	beneficiaryRepo BeneficiaryRepository
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUser interface {
	GetID() int
	GetByEmail() string
	GetPassword() string
}
