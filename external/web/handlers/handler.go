package handlers

import (
	"context"

	"github.com/samuelralmeida/neofarma/internal/patient"
	"github.com/samuelralmeida/neofarma/internal/user"
)

type PatientUseCases interface {
	Save(ctx context.Context, input *patient.NewPatientInputDto) (*patient.PatientOutputDto, error)
	GetById(ctx context.Context, id string) (*patient.PatientOutputDto, error)
}

type UserUseCases interface {
	Create(ctx context.Context, input *user.CreateUserInputDto) (*user.UserOutputDto, error)
	Authenticate(ctx context.Context, email, password string) (*user.User, error)
}

type WebHandler struct {
	PatientUseCases
	UserUseCases
}

func NewWebHandler(patientUseCases PatientUseCases, userUseCases UserUseCases) *WebHandler {
	return &WebHandler{PatientUseCases: patientUseCases, UserUseCases: userUseCases}
}
