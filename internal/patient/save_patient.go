package patient

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/user"
)

func (s *PatientUseCases) Save(ctx context.Context, input *NewPatientInputDto) (*PatientOutputDto, error) {
	_, err := s.userUseCases.LoggedUserPermission(ctx, user.AdminHierarchy)
	if err != nil {
		return nil, fmt.Errorf("error to authorize logged user to create new user: %w", err)
	}

	patient := Patient{
		Name:  input.Name,
		Cpf:   input.Cpf,
		Email: input.Email,
		Phone: input.Phone,
	}

	err = s.patientRepository.SavePatient(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return &PatientOutputDto{
		ID:    patient.ID,
		Name:  patient.Name,
		Cpf:   patient.Cpf,
		Email: patient.Email,
		Phone: patient.Phone,
	}, nil
}
