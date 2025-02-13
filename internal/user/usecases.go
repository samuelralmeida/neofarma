package user

import (
	"context"
)

type UserRepository interface {
	SaveUser(context.Context, *User) error
	GetUsersByEmail(ctx context.Context, email string) ([]User, error)
	UpdatePassword(ctx context.Context, userID, newPasswordHash string) error
	GetUserByID(ctx context.Context, id string) (*User, error)
}

type UserUseCases struct {
	userRepository UserRepository
}

func NewUserUseCases(userRepository UserRepository) *UserUseCases {
	return &UserUseCases{userRepository: userRepository}
}
