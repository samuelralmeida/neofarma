package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/responsibility"
)

func (ruc *ResponsibilityUseCases) ListPatientsByUser(ctx context.Context, userID string) ([]responsibility.PatientWithRelationship, error) {
	isUserExists, err := ruc.userUseCases.CheckUserExists(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error to check if user exists: %w", err)
	}
	if !isUserExists {
		return nil, errors.New("user nof found")
	}

	result, err := ruc.responsibilityRepository.GetPatientsByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error to get patients by user: %w", err)
	}
	return result, nil
}
