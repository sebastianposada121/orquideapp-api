package employee_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Create(employee domain.Employee) error {
	query := "INSERT INTO employees (name, last_name, email, document_id, document_number, phone, ips_id, rol_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, employee.Name, employee.LastName, employee.Email, employee.DocumentId, employee.DocumentNumber, employee.Phone, employee.IpsId, employee.RolId)
	return err
}

func (r *Repository) CreateRol(rol domain.Rol) error {
	query := "INSERT INTO roles (name, description) VALUES (?, ?)"
	_, err := r.DB.Exec(query, rol.Name, rol.Description)
	return err
}

func (r *Repository) GetByEmail(email string) (domain.Employee, error) {
	var employee domain.Employee
	query := "SELECT id, name, email, ips_id, password FROM employees WHERE email = ?"
	row := r.DB.QueryRow(query, email)
	err := row.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.IpsId, &employee.Password)
	return employee, err
}

func (r *Repository) UpdatePassword(credentials domain.Login) error {
	query := "UPDATE employees SET password = ? where email = ?"
	_, err := r.DB.Exec(query, credentials.Password, credentials.Email)
	return err
}

func (r *Repository) Delete(id int) error {
	return nil
}
