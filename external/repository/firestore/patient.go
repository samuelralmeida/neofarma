package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/samuelralmeida/neofarma/internal/patient"
)

type FirestoreRepository struct {
	client *firestore.Client
}

func NewFirestoreRepository(client *firestore.Client) *FirestoreRepository {
	return &FirestoreRepository{client: client}
}

func (f *FirestoreRepository) Save(ctx context.Context, patient *patient.Patient) error {
	documentRef, _, err := f.client.Collection("patient").Add(ctx, map[string]interface{}{
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

func (f *FirestoreRepository) GetById(ctx context.Context, id string) (*patient.Patient, error) {
	doc, err := f.client.Collection("patient").Doc(id).Get(ctx)
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
