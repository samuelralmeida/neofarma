package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCases) UpdatePassword(ctx context.Context, userID, password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error to generate password hash: %w", err)
	}

	passwordHash := string(hashedBytes)
	err = u.userRepository.UpdatePassword(ctx, userID, passwordHash)
	if err != nil {
		return fmt.Errorf("error to update password: %w", err)
	}

	return nil
}
