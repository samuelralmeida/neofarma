package handlers

import (
	"encoding/json"
	"net/http"
)

// @Summary		Criar uma relação entre usuário e paciente
// @Description	Recebe os dados de paciente, usuário e tipo de relação
// @Tags			Responsibility
// @Accept			json
// @Produce		json
// @Param			userId		body		string	true	"user id"
// @Param			patientId	body		string	true	"patient id"
// @Param			bond	body		string	true	"relationship type"
// @Success		200		{string}	string	"relação criada com sucesso"
// @Failure		400		{string}	string						"Erro de validação dos dados"
// @Failure		500		{string}	string						"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/responsibilities/save [post]
func (wh *WebHandler) CreateRelationship(w http.ResponseWriter, r *http.Request) {
	input := struct {
		UserID           string `json:"userId"`
		PatientID        string `json:"patientId"`
		RelationshipType string `json:"bond"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = wh.ResponsibilityUseCases.LinkUserToPatient(r.Context(), input.UserID, input.PatientID, input.RelationshipType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary		Remove relação entre usuário e paciente
// @Description	Recebe os dados de paciente, usuário e tipo de relação
// @Tags			Responsibility
// @Accept			json
// @Produce		json
// @Param			userId		body		string	true	"user id"
// @Param			patientId	body		string	true	"patient id"
// @Param			bond	body		string	true	"relationship type"
// @Success		200		{string}	string	"Relação removida com sucesso"
// @Failure		400		{string}	string						"Erro de validação dos dados"
// @Failure		500		{string}	string						"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/responsibilities/remove [post]
func (wh *WebHandler) RemoveRelationship(w http.ResponseWriter, r *http.Request) {
	input := struct {
		UserID           string `json:"userId"`
		PatientID        string `json:"patientId"`
		RelationshipType string `json:"bond"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = wh.ResponsibilityUseCases.UnlinkUserFromPatient(r.Context(), input.UserID, input.PatientID, input.RelationshipType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
