package history

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"

	"github.com/hrvadl/book-service/internal/domain/history"
)

const collection = "history"

func NewRepo(db *firestore.Client) *Repo {
	return &Repo{
		collection: collection,
		db:         db,
	}
}

type Repo struct {
	collection string
	db         *firestore.Client
}

func (r *Repo) Add(ctx context.Context, h history.ReadingHistory) (string, error) {
	ref, _, err := r.db.Collection(r.collection).Add(ctx, h)
	if err != nil {
		return "", fmt.Errorf("failed to add reading history: %w", err)
	}

	return ref.ID, nil
}

func (r *Repo) GetByUserID(ctx context.Context, userID int) ([]history.ReadingHistory, error) {
	docs, err := r.db.Collection(r.collection).
		Where("user_id", "==", userID).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %w", err)
	}

	return docToHistory(docs)
}

func docToHistory(docs []*firestore.DocumentSnapshot) ([]history.ReadingHistory, error) {
	histories := make([]history.ReadingHistory, 0, len(docs))
	for _, doc := range docs {
		var h history.ReadingHistory
		if err := doc.DataTo(&h); err != nil {
			return nil, fmt.Errorf("failed to unmarshall data: %w", err)
		}
		histories = append(histories, h)
	}

	return histories, nil
}
