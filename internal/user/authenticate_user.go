package user

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCases) Authenticate(ctx context.Context, email, password string) (*User, error) {
	user, err := u.getByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error to find user: %w", err)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password invalid: %w", err)
	}

	return user, nil
}
