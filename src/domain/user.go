package domain

type UserRepository interface {
	GetByEmail(email string) (User, error)
	GetByID(id int) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(id int) error
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nit      string `json:"nit"`
	Phone    string `json:"phone"`
}
