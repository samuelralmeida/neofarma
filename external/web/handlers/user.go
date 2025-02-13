package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samuelralmeida/neofarma/internal/auth"
	"github.com/samuelralmeida/neofarma/internal/user"
)

func (wh *WebHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input user.CreateUserInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := wh.UserUseCases.Create(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = auth.SetCookie(w, output.ID)
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

func (wh *WebHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := wh.UserUseCases.Authenticate(r.Context(), input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = auth.SetCookie(w, output.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (wh *WebHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	auth.DeleteCookie(w)
}
