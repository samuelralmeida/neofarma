package patient

import "context"

type PatientRepository interface {
	SavePatient(context.Context, *Patient) error
	GetPatientById(ctx context.Context, id string) (*Patient, error)
}

type PatientUseCases struct {
	patientRepository PatientRepository
}

func NewPatientUseCases(patientRepository PatientRepository) *PatientUseCases {
	return &PatientUseCases{patientRepository: patientRepository}
}
