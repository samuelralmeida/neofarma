package firestore

import (
	"context"
	"fmt"

	"github.com/samuelralmeida/neofarma/internal/patient"
)

func (f *FirestoreRepository) SavePatient(ctx context.Context, patient *patient.Patient) error {
	documentRef, _, err := f.client.Collection("patients").Add(ctx, map[string]interface{}{
		"cpf":   patient.Cpf,
		"email": patient.Email,
		"name":  patient.Name,
		"phone": patient.Phone,
	})
	if err != nil {
		return fmt.Errorf("error to add patient: %w", err)
	}

	patient.ID = documentRef.ID
	return nil
}

func (f *FirestoreRepository) GetPatientById(ctx context.Context, id string) (*patient.Patient, error) {
	doc, err := f.client.Collection("patients").Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error to get patient by id: %w", err)
	}
	var p patient.Patient
	err = doc.DataTo(&p)
	if err != nil {
		return nil, fmt.Errorf("error to extract patient document: %w", err)
	}

	p.ID = doc.Ref.ID
	return &p, nil
}
