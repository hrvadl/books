package history

import "time"

type ReadingHistory struct {
	LastOpened time.Time `firestore:"last_opened"`
	BookID     int       `firestore:"book_id"`
	UserID     int       `firestore:"user_id"`
}
