package patient

type NewPatientInputDto struct {
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type PatientOutputDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
