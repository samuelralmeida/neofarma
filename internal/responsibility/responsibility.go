package responsibility

import (
	"github.com/samuelralmeida/neofarma/internal/patient"
	"github.com/samuelralmeida/neofarma/internal/user"
)

type PatientWithRelationship struct {
	Patient          patient.Patient
	RelationshipType string
	UserID           string
}

type UserWithRelationship struct {
	User             user.User
	RelationshipType string
	PatientID        string
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
