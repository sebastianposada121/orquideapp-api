package appointmentstep_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Create(appointmentStep domain.AppointmentStep) error {
	query := "INSERT INTO appointment_steps (medical_appointment_id, name, status) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, appointmentStep.MedicalAppointmentID, appointmentStep.Name, appointmentStep.Status)
	return err
}
func (r *Repository) Update(appointmentStep domain.AppointmentStep) error {
	return nil
}
func (r *Repository) Delete(id int) error {
	return nil
}

func (r *Repository) GetAllByMedicalAppointmentID(id int) (domain.AppointmentSteps, error) {
	return nil, nil
}
