package history

import "time"

type ReadingHistory struct {
	LastOpened time.Time
	BookID     int
	UserID     int
}
