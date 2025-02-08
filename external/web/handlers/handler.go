package handlers

import (
	"context"

	"github.com/samuelralmeida/neofarma/internal/patient"
)

type PatientUseCases interface {
	Save(ctx context.Context, input *patient.NewPatientInputDto) (*patient.PatientOutputDto, error)
	GetById(ctx context.Context, id string) (*patient.PatientOutputDto, error)
}

type WebHandler struct {
	PatientUseCases
}

func NewWebHandler(patientUseCases PatientUseCases) *WebHandler {
	return &WebHandler{PatientUseCases: patientUseCases}
}
