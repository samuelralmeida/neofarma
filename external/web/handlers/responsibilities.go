package handlers

import (
	"encoding/json"
	"net/http"
)

// @Summary		Criar uma relação entre usuário e paciente
// @Description	Recebe os dados de paciente, usuário e tipo de relação
// @Tags			Relationship
// @Accept			json
// @Produce		json
// @Param			userId		body		string	true	"user id"
// @Param			patientId	body		string	true	"patient id"
// @Param			bond		body		string	true	"relationship type"
// @Success		200			{string}	string	"relação criada com sucesso"
// @Failure		400			{string}	string	"Erro de validação dos dados"
// @Failure		500			{string}	string	"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/relationships/save [post]
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
// @Tags			Relationship
// @Accept			json
// @Produce		json
// @Param			userId		body		string	true	"user id"
// @Param			patientId	body		string	true	"patient id"
// @Param			bond		body		string	true	"relationship type"
// @Success		200			{string}	string	"Relação removida com sucesso"
// @Failure		400			{string}	string	"Erro de validação dos dados"
// @Failure		500			{string}	string	"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/relationships/remove [post]
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

// @Summary		Lista os usuários com relação ao paciente
// @Description	Devolve os usuário que tem relação com o paciente informado
// @Tags			Relationship
// @Accept			json
// @Produce		json
// @Param			id	path		string									true	"patient id"
// @Success		200			{object}	[]responsibility.UserWithRelationship	"Usuários com relação com o paciente informado"
// @Failure		500			{string}	string									"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/patients/{id}/users/relationships [get]
func (wh *WebHandler) ListUsersByPatient(w http.ResponseWriter, r *http.Request) {
	patientID := r.PathValue("id")

	output, err := wh.ResponsibilityUseCases.ListUsersByPatient(r.Context(), patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// @Summary		Lista os usuários com relação ao paciente
// @Description	Devolve os usuário que tem relação com o paciente informado
// @Tags			Relationship
// @Accept			json
// @Produce		json
// @Param			id	path		string									true	"user id"
// @Success		200			{object}	[]responsibility.PatientWithRelationship	"Pacientes com relação com o usuário informado"
// @Failure		500			{string}	string									"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/users/{id}/patients/relationships [get]
func (wh *WebHandler) ListPatientsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")

	output, err := wh.ResponsibilityUseCases.ListPatientsByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
