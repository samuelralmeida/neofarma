package usecases

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/responsibility"
	"github.com/samuelralmeida/neofarma/internal/user"
)

func (ruc *ResponsibilityUseCases) UnlinkUserFromPatient(ctx context.Context, userID, patientID, relationshipType string) error {
	_, err := ruc.userUseCases.LoggedUserPermission(ctx, user.AdminHierarchy)
	if err != nil {
		return fmt.Errorf("error to authorize logged user to remove relationship: %w", err)
	}

	_, ok := responsibility.GetRelationshipByName(relationshipType)
	if !ok {
		return fmt.Errorf("invalid relationship: %s", relationshipType)
	}

	err = ruc.responsibilityRepository.DeleteRelationship(ctx, userID, patientID, relationshipType)
	if err != nil {
		return fmt.Errorf("error to delete relationship: %w", err)
	}
	return nil
}
