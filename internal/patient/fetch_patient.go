package patient

import (
	"context"
	"fmt"
)

func (s *PatientUseCases) GetById(ctx context.Context, id string) (*PatientOutputDto, error) {
	patient, err := s.patientRepository.GetPatientById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error to get patient by id: %w", err)
	}

	return &PatientOutputDto{
		ID:    patient.ID,
		Name:  patient.Name,
		Cpf:   patient.Cpf,
		Email: patient.Email,
		Phone: patient.Phone,
	}, nil
}
