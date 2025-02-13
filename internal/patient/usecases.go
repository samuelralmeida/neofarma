package patient

import (
	"context"

	"github.com/samuelralmeida/neofarma/internal/user"
)

type PatientRepository interface {
	SavePatient(context.Context, *Patient) error
	GetPatientById(context.Context, string) (*Patient, error)
}

type UserUseCases interface {
	LoggedUserPermission(context.Context, user.HierarchyRole) (*user.User, error)
}

type PatientUseCases struct {
	patientRepository PatientRepository
	userUseCases      UserUseCases
}

func NewPatientUseCases(patientRepository PatientRepository, userUseCases UserUseCases) *PatientUseCases {
	return &PatientUseCases{patientRepository: patientRepository, userUseCases: userUseCases}
}
