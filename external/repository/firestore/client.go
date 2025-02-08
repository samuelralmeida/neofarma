package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClientWithDatabase(ctx, os.Getenv("PROJECT_ID"), os.Getenv("FIRESTORE"))
	if err != nil {
		log.Fatalln(err)
	}

	return client, nil
}
