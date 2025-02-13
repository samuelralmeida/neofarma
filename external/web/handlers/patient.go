package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samuelralmeida/neofarma/internal/patient"
)

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
