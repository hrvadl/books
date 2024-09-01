package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewFirestore(ctx context.Context, filename, projectID, db string) (*firestore.Client, error) {
	sa := option.WithCredentialsFile(filename)
	client, err := firestore.NewClientWithDatabase(ctx, projectID, db, sa)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase app: %w", err)
	}

	return client, nil
}
