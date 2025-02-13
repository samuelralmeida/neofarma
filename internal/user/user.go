package user

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Role         Role   `json:"role"`
	Origin       Origin `json:"origin"`
}
