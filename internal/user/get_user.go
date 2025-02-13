package user

import (
	"context"
	"fmt"
)

func (u *UserUseCases) getByEmail(ctx context.Context, email string) (*User, error) {
	users, err := u.userRepository.GetUsersByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error to get user by email: %w", err)
	}

	if len(users) == 0 {
		return nil, nil
	}

	if len(users) > 1 {
		return nil, fmt.Errorf("duplicated user: %w", err)
	}

	return &users[0], nil
}

func (u *UserUseCases) getByID(ctx context.Context, id string) (*User, error) {
	user, err := u.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error to get user by id: %w", err)
	}
	return user, nil
}
