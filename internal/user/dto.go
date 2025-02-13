package user

type CreateUserInputDto struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Origin   string `json:"origin"`
}

type UserOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
