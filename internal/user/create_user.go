package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCases) Create(ctx context.Context, input *CreateUserInputDto) (*UserOutputDto, error) {
	_, err := u.LoggedUserPermission(ctx, AdminHierarchy)
	if err != nil {
		return nil, fmt.Errorf("error to authorize logged user to create new user: %w", err)
	}

	existedUser, err := u.getByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("error to check if user already exists: %w", err)
	}

	if existedUser != nil {
		return nil, fmt.Errorf("user already exists: %w", err)
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("generate password hash: %w", err)
	}

	passwordHash := string(hashedBytes)

	role, ok := GetRoleByName(input.Role)
	if !ok {
		return nil, errors.New("role not found")
	}

	origin, ok := GetOriginByName(input.Origin)
	if !ok {
		return nil, errors.New("origin not found")
	}

	user := User{
		Email:        strings.ToLower(input.Email),
		PasswordHash: passwordHash,
		Role:         role,
		Origin:       origin,
	}

	err = u.userRepository.SaveUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &UserOutputDto{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role.Name,
	}, nil
}
