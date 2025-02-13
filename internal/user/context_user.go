package user

import (
	"context"
	"errors"
	"log"
)

type contextKey string

const loggedUserKey contextKey = "logged-user"

func (u *UserUseCases) SetUserToContext(ctx context.Context, userId string) (context.Context, error) {
	user, err := u.getByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, loggedUserKey, user), nil
}

func (u *UserUseCases) getUserFromContext(ctx context.Context) *User {
	user, _ := ctx.Value(loggedUserKey).(*User)
	return user
}

func (u *UserUseCases) LoggedUsedPermission(ctx context.Context, hierarchy HierarchyRole) (*User, error) {
	user := u.getUserFromContext(ctx)
	if user == nil {
		return nil, errors.New("user not logged")
	}

	if !user.Role.HasHierarchy(hierarchy) {
		log.Printf("userID: %s with role %d is trying to do something that required role %d\n", user.ID, user.Role.Hierarchy, hierarchy)
		return nil, errors.New("user unauthorized")
	}

	return user, nil
}
