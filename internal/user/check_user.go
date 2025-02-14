package user

import (
	"context"
	"fmt"
)

func (u *UserUseCases) CheckUserExists(ctx context.Context, userID string) (bool, error) {
	user, err := u.getByID(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("error to find user: %w", err)
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}
