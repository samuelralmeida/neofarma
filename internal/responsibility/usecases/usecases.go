package usecases

import (
	"context"

	"github.com/samuelralmeida/neofarma/internal/responsibility"
	"github.com/samuelralmeida/neofarma/internal/user"
)

type ResponsibilityRepository interface {
	CreateRelationship(ctx context.Context, userID, patientID, relationshipType string) error
	DeleteRelationship(ctx context.Context, userID, patientID, relationshipType string) error
	GetPatientsByUser(ctx context.Context, userID string) ([]responsibility.PatientWithRelationship, error)
	GetUsersByPatient(ctx context.Context, patientID string) ([]responsibility.UserWithRelationship, error)
	ExistsRelationship(ctx context.Context, userID, patientID, relationshipType string) (bool, error)
}

type UserUseCases interface {
	CheckUserExists(ctx context.Context, userID string) (bool, error)
	LoggedUserPermission(context.Context, user.HierarchyRole) (*user.User, error)
}

type PatientUseCases interface {
	CheckPatientExists(ctx context.Context, patientID string) (bool, error)
}

type ResponsibilityUseCases struct {
	responsibilityRepository ResponsibilityRepository
	userUseCases             UserUseCases
	patientUseCases          PatientUseCases
}

func NewResponsibilityUseCases(responsibilityRepository ResponsibilityRepository, userUseCases UserUseCases, patientUseCases PatientUseCases) *ResponsibilityUseCases {
	return &ResponsibilityUseCases{responsibilityRepository: responsibilityRepository, userUseCases: userUseCases, patientUseCases: patientUseCases}
}
