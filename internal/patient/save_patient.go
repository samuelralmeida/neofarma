package patient

import (
	"context"
)

func (s *PatientUseCases) Save(ctx context.Context, input *NewPatientInputDto) (*PatientOutputDto, error) {
	patient := Patient{
		Name:  input.Name,
		Cpf:   input.Cpf,
		Email: input.Email,
		Phone: input.Phone,
	}

	err := s.patientRepository.Save(ctx, &patient)
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
