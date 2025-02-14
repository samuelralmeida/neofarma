package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/samuelralmeida/neofarma/internal/responsibility"
	"google.golang.org/api/iterator"
)

const responsibilitiesCollectionName = "responsibilities"

type storeResponsibility struct {
	PatientID        string `firestore:"patientId"`
	UserID           string `firestore:"userId"`
	RelationshipType string `firestore:"relationshipType"`
}

func (f *FirestoreRepository) CreateRelationship(ctx context.Context, userID, patientID, relationshipType string) error {
	sr := storeResponsibility{
		PatientID:        patientID,
		UserID:           userID,
		RelationshipType: relationshipType,
	}

	_, _, err := f.client.Collection(responsibilitiesCollectionName).Add(ctx, sr)
	if err != nil {
		return fmt.Errorf("error to add relationship: %w", err)
	}
	return nil
}

func (f *FirestoreRepository) DeleteRelationship(ctx context.Context, userID, patientID, relationshipType string) error {
	err := f.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		iter := f.client.Collection(responsibilitiesCollectionName).Where("userId", "==", userID).Where("patientId", "==", patientID).Where("relationshipType", "==", relationshipType).Documents(ctx)
		defer iter.Stop()

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return fmt.Errorf("error to iterate responsibilities: %w", err)
			}

			err = tx.Delete(doc.Ref)
			if err != nil {
				return fmt.Errorf("error to delete responsibilities: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error to run transaction to delete responsibilities")
	}

	return nil
}
func (f *FirestoreRepository) GetPatientsByUser(ctx context.Context, userID string) ([]responsibility.PatientWithRelationship, error) {
	iter := f.client.Collection(responsibilitiesCollectionName).Where("userId", "==", userID).Documents(ctx)
	defer iter.Stop()

	responsibilities := []storeResponsibility{}
	patientIds := []string{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error to get patients by user: %w", err)
		}

		var s storeResponsibility
		err = doc.DataTo(&s)
		if err != nil {
			return nil, fmt.Errorf("error to populate user: %w", err)
		}
		responsibilities = append(responsibilities, s)
		patientIds = append(patientIds, s.PatientID)
	}

	iterP := f.client.Collection(patientsCollectionName).Where("id", "in", patientIds).Documents(ctx)
	defer iterP.Stop()

	patients := map[string]storePatient{}
	for {
		doc, err := iterP.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error to get patients by user: %w", err)
		}

		var s storePatient
		err = doc.DataTo(&s)
		if err != nil {
			return nil, fmt.Errorf("error to populate user: %w", err)
		}
		patients[doc.Ref.ID] = s
	}

	result := make([]responsibility.PatientWithRelationship, len(responsibilities))
	for i, r := range responsibilities {
		p := patients[r.PatientID]
		result[i] = responsibility.PatientWithRelationship{
			Patient:          p.toPatient(r.PatientID),
			RelationshipType: r.RelationshipType,
			UserID:           r.UserID,
		}
	}

	return result, nil
}
func (f *FirestoreRepository) GetUsersByPatient(ctx context.Context, patientID string) ([]responsibility.UserWithRelationship, error) {
	iter := f.client.Collection(responsibilitiesCollectionName).Where("patientId", "==", patientID).Documents(ctx)
	defer iter.Stop()

	responsibilities := []storeResponsibility{}
	userIds := []string{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error to get patients by user: %w", err)
		}

		var s storeResponsibility
		err = doc.DataTo(&s)
		if err != nil {
			return nil, fmt.Errorf("error to populate user: %w", err)
		}
		responsibilities = append(responsibilities, s)
		userIds = append(userIds, s.PatientID)
	}

	iterP := f.client.Collection(userCollectionName).Where("id", "in", userIds).Documents(ctx)
	defer iterP.Stop()

	users := map[string]storeUser{}
	for {
		doc, err := iterP.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error to get patients by user: %w", err)
		}

		var s storeUser
		err = doc.DataTo(&s)
		if err != nil {
			return nil, fmt.Errorf("error to populate user: %w", err)
		}
		users[doc.Ref.ID] = s
	}

	result := make([]responsibility.UserWithRelationship, len(responsibilities))
	for i, r := range responsibilities {
		u := users[r.UserID]
		result[i] = responsibility.UserWithRelationship{
			User:             u.toUser(r.UserID),
			RelationshipType: r.RelationshipType,
			PatientID:        r.PatientID,
		}
	}

	return result, nil

}
func (f *FirestoreRepository) ExistsRelationship(ctx context.Context, userID, patientID, relationshipType string) (bool, error) {
	panic("not implemented")

}
