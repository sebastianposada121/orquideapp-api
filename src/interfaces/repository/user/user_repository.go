package user_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetByID(id int) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (r *Repository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := r.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (r *Repository) Create(user domain.User) error {
	query := "INSERT INTO users (name, email, password, nit, phone) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.Nit, user.Phone)
	return err
}

func (r *Repository) Update(user domain.User) error {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}
