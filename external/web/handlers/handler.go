package handlers

import (
	"context"

	"github.com/samuelralmeida/neofarma/internal/patient"
	"github.com/samuelralmeida/neofarma/internal/responsibility"
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

type ResponsibilityUseCases interface {
	LinkUserToPatient(ctx context.Context, userID, patientID, relationshipType string) error
	UnlinkUserFromPatient(ctx context.Context, userID, patientID, relationshipType string) error
	ListUsersByPatient(ctx context.Context, patientID string) ([]responsibility.UserWithRelationship, error)
	ListPatientsByUser(ctx context.Context, userID string) ([]responsibility.PatientWithRelationship, error)
}

type WebHandler struct {
	PatientUseCases
	UserUseCases
	ResponsibilityUseCases
}

func NewWebHandler(patientUseCases PatientUseCases, userUseCases UserUseCases, responsibilityUseCases ResponsibilityUseCases) *WebHandler {
	return &WebHandler{PatientUseCases: patientUseCases, UserUseCases: userUseCases, ResponsibilityUseCases: responsibilityUseCases}
}
