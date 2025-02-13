package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samuelralmeida/neofarma/internal/auth"
	"github.com/samuelralmeida/neofarma/internal/user"
)

// @Summary		Cria um novo usuário
// @Description	Cria um novo usuário com os dados fornecidos e retorna os detalhes do usuário, incluindo um cookie de autenticação.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			input	body		user.CreateUserInputDto	true	"Dados do novo usuário"
// @Success		200		{object}	user.UserOutputDto		"Usuário criado com sucesso"
// @Failure		400		{string}	string					"Erro de validação dos dados"
// @Failure		500		{string}	string					"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/admin/create [post]
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

// @Summary		Autentica um usuário e retorna um cookie de autenticação
// @Description	Autentica um usuário com email e senha fornecidos e retorna um cookie para manter a sessão do usuário.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			email		body		string	true	"user email"
// @Param			password	body		string	true	"user password"
// @Success		200			{string}	string	"Autenticação bem-sucedida"
// @Failure		400			{string}	string	"Erro de validação dos dados"
// @Failure		500			{string}	string	"Erro interno do servidor"
// @Security ApiKeyAuth
// @Router			/users/signin [post]
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

// @Summary		Desconecta o usuário removendo o cookie de autenticação
// @Description	Remove o cookie de autenticação do usuário, invalidando a sessão.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"Desconexão bem-sucedida"
// @Failure		500	{string}	string	"Erro interno do servidor"
// @Security		ApiKeyAuth
// @Router			/users/signout [post]
func (wh *WebHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	auth.DeleteCookie(w)
}
