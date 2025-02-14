package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/responsibility"
	"github.com/samuelralmeida/neofarma/internal/user"
)

func (ruc *ResponsibilityUseCases) LinkUserToPatient(ctx context.Context, userID, patientID, relationshipType string) error {
	_, err := ruc.userUseCases.LoggedUserPermission(ctx, user.AdminHierarchy)
	if err != nil {
		return fmt.Errorf("error to authorize logged user to create relationship: %w", err)
	}

	isUserExists, err := ruc.userUseCases.CheckUserExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("error to check if user exists: %w", err)
	}
	if !isUserExists {
		return errors.New("user nof found")
	}

	isPatientExists, err := ruc.patientUseCases.CheckPatientExists(ctx, patientID)
	if err != nil {
		return fmt.Errorf("error to check if patient exists: %w", err)
	}
	if !isPatientExists {
		return errors.New("patient nof found")
	}

	_, ok := responsibility.GetRelationshipByName(relationshipType)
	if !ok {
		return fmt.Errorf("invalid relationship: %s", relationshipType)
	}

	err = ruc.responsibilityRepository.CreateRelationship(ctx, userID, patientID, relationshipType)
	if err != nil {
		return fmt.Errorf("error to create relationship: %w", err)
	}
	return nil
}
