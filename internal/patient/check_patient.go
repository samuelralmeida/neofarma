package patient

import (
	"context"
	"fmt"
)

func (s *PatientUseCases) CheckPatientExists(ctx context.Context, patientID string) (bool, error) {
	patient, err := s.GetById(ctx, patientID)
	if err != nil {
		return false, fmt.Errorf("error to find patient: %w", err)
	}

	if patient == nil {
		return false, nil
	}

	return true, nil
}
