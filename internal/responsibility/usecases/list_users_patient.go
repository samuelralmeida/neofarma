package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/responsibility"
)

func (ruc *ResponsibilityUseCases) ListUsersByPatient(ctx context.Context, patientID string) ([]responsibility.UserWithRelationship, error) {
	isUserExists, err := ruc.patientUseCases.CheckPatientExists(ctx, patientID)
	if err != nil {
		return nil, fmt.Errorf("error to check if user exists: %w", err)
	}
	if !isUserExists {
		return nil, errors.New("user nof found")
	}

	result, err := ruc.responsibilityRepository.GetUsersByPatient(ctx, patientID)
	if err != nil {
		return nil, fmt.Errorf("error to get patients by user: %w", err)
	}
	return result, nil
}
