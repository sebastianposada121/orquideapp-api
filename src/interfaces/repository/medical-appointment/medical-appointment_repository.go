package medicalAppointment_repository

import (
	"database/sql"
	"orquideapp/src/domain"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateMedicalAppointment(medicalEndpoint domain.MedicalAppointment) (int, error) {
	query := "INSERT INTO medical_appointments (beneficiary_id, appointment_date, service_id, description) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, medicalEndpoint.BeneficiaryID, medicalEndpoint.AppointmentDate, medicalEndpoint.ServiceID, medicalEndpoint.Description)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return int(id), err
	}

	return int(id), err
}

func (r *Repository) GetAllMedicalAppointmentByBeneficiaryID(id int) (domain.MedicalAppointments, error) {
	medicalAppointmens := domain.MedicalAppointments{}
	query := "SELECT id, active, description, service_id, appointment_date FROM medical_appointments WHERE beneficiary_id = ?"
	data, err := r.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	for data.Next() {
		medicalAppointment := domain.MedicalAppointment{}
		data.Scan(&medicalAppointment.ID, &medicalAppointment.Active, &medicalAppointment.Description, &medicalAppointment.ServiceID, &medicalAppointment.AppointmentDate)
		medicalAppointmens = append(medicalAppointmens, medicalAppointment)
	}

	return medicalAppointmens, nil
}

func (r *Repository) CreateAppointmentStep(appointmentStep domain.AppointmentStep) error {
	query := "INSERT INTO appointment_steps (medical_appointment_id, name, status) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, appointmentStep.MedicalAppointmentID, appointmentStep.Name, appointmentStep.Status)
	return err
}

func (r *Repository) GetAllAppointmentStepsByMedicalAppointmentID(id int) (domain.AppointmentSteps, error) {
	appointmentSteps := domain.AppointmentSteps{}
	query := "SELECT id, name, description, status, created_at FROM appointment_steps WHERE medical_appointment_id = ?"
	data, err := r.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	for data.Next() {
		appointmentStep := domain.AppointmentStep{}
		data.Scan(&appointmentStep.ID, &appointmentStep.Name, &appointmentStep.Description, &appointmentStep.Status, &appointmentStep.CreatedAt)
		appointmentSteps = append(appointmentSteps, appointmentStep)
	}

	return appointmentSteps, nil
}
