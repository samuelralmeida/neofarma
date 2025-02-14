package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samuelralmeida/neofarma/internal/patient"
)

//	@Summary		Salva um novo paciente
//	@Description	Recebe os dados de um novo paciente e os salva no sistema.
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			input	body		patient.NewPatientInputDto	true	"Dados do paciente"
//	@Success		200		{object}	patient.PatientOutputDto	"Paciente salvo com sucesso"
//	@Failure		400		{string}	string						"Erro de validação dos dados"
//	@Failure		500		{string}	string						"Erro interno do servidor"
//	@Security		ApiKeyAuth
//	@Router			/patients/save [post]
func (wh *WebHandler) SavePatient(w http.ResponseWriter, r *http.Request) {
	var input patient.NewPatientInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := wh.PatientUseCases.Save(r.Context(), &input)
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

//	@Summary		Lista os usuários com relação ao paciente
//	@Description	Devolve os usuário que tem relação com o paciente informado
//	@Tags			Relationship
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"user id"
//	@Success		200	{object}	patient.PatientOutputDto	"Dados do paciente"
//	@Failure		500	{string}	string						"Erro interno do servidor"
//	@Security		ApiKeyAuth
//	@Router			/patients/{id} [get]
func (wh *WebHandler) GetPatientById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	output, err := wh.PatientUseCases.GetById(r.Context(), id)
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
