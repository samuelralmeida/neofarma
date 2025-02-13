package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/samuelralmeida/neofarma/internal/user"
	"google.golang.org/api/iterator"
)

const userCollectionName = "users"

type storeUser struct {
	Email    string `firestore:"email"`
	Password string `firestore:"password"`
	Role     string `firestore:"role"`
	Origin   string `firestore:"origin"`
}

func newStoreUser(u *user.User) storeUser {
	return storeUser{
		Email:    u.Email,
		Password: u.PasswordHash,
		Role:     u.Role.Name,
		Origin:   string(u.Origin),
	}
}

func (su *storeUser) toUser(id string) user.User {
	role, _ := user.GetRoleByName(su.Role)
	origin, _ := user.GetOriginByName(su.Origin)

	return user.User{
		ID:           id,
		Email:        su.Email,
		PasswordHash: su.Password,
		Role:         role,
		Origin:       origin,
	}
}

func (f *FirestoreRepository) SaveUser(ctx context.Context, user *user.User) error {
	userStore := newStoreUser(user)
	documentRef, _, err := f.client.Collection(userCollectionName).Add(ctx, userStore)
	if err != nil {
		return fmt.Errorf("error to add patient: %w", err)
	}

	user.ID = documentRef.ID
	return nil
}

func (f *FirestoreRepository) GetUsersByEmail(ctx context.Context, email string) ([]user.User, error) {
	iter := f.client.Collection(userCollectionName).Where("email", "==", email).Documents(ctx)
	defer iter.Stop()

	users := []user.User{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error to get user by email: %w", err)
		}

		var u storeUser
		err = doc.DataTo(&u)
		if err != nil {
			return nil, fmt.Errorf("error to populate user: %w", err)
		}
		users = append(users, u.toUser(doc.Ref.ID))
	}

	return users, nil
}

func (f *FirestoreRepository) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	doc, err := f.client.Collection(userCollectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error to get patient by id: %w", err)
	}
	var u storeUser
	err = doc.DataTo(&u)
	if err != nil {
		return nil, fmt.Errorf("error to extract patient document: %w", err)
	}

	user := u.toUser(doc.Ref.ID)
	return &user, nil
}

func (f *FirestoreRepository) UpdatePassword(ctx context.Context, id, newPasswordHash string) error {
	_, err := f.client.Collection("users").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "password",
			Value: newPasswordHash,
		},
	})
	if err != nil {
		return fmt.Errorf("error to update password in firestore: %w", err)
	}
	return nil
}
