package patient

import "context"

type PatientRepository interface {
	Save(context.Context, *Patient) error
	GetById(ctx context.Context, id string) (*Patient, error)
}

type PatientUseCases struct {
	patientRepository PatientRepository
}

func NewPatientUseCases(patientRepository PatientRepository) *PatientUseCases {
	return &PatientUseCases{patientRepository: patientRepository}
}
