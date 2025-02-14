package responsibility

import (
	"github.com/samuelralmeida/neofarma/internal/patient"
	"github.com/samuelralmeida/neofarma/internal/user"
)

type PatientWithRelationship struct {
	Patient          patient.Patient `json:"patient"`
	RelationshipType string          `json:"relationshipType"`
	UserID           string          `json:"userId"`
}

type UserWithRelationship struct {
	User             user.User `json:"user"`
	RelationshipType string    `json:"relationshipType"`
	PatientID        string    `json:"patientId"`
}

type RelationshipType string

var (
	Financial RelationshipType = "financial"
	Clinical  RelationshipType = "clinical"
	B2B       RelationshipType = "b2b"
	B2B2C     RelationshipType = "b2b2c"
)

var relationshipTypes = map[string]RelationshipType{
	string(Financial): Financial,
	string(Clinical):  Clinical,
	string(B2B):       B2B,
	string(B2B2C):     B2B2C,
}

func GetRelationshipByName(relationship string) (RelationshipType, bool) {
	o, exists := relationshipTypes[relationship]
	return o, exists
}
