package domain

type EmployeesRepository interface {
	// Login(credentials Login) error
	Create(employee Employee) error
	Delete(id int) error
	CreateRol(rol Rol) error
	GetByEmail(email string) (Employee, error)
	UpdatePassword(credentials Login) error
}

type Employee struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	DocumentId     int    `json:"document_id"`
	DocumentNumber string `json:"document_number"`
	Phone          string `json:"phone"`
	IpsId          int    `json:"ips_id"`
	RolId          int    `json:"rol_id"`
}
