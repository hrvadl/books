package recommendation

import (
	"github.com/hrvadl/book-service/internal/domain/book"
	"github.com/hrvadl/book-service/internal/domain/history"
)

func newHistoryComparator(h []history.ReadingHistory) *historyComparator {
	return &historyComparator{
		history: h,
	}
}

type historyComparator struct {
	history []history.ReadingHistory
}

func (hc *historyComparator) Contains(b book.Book) bool {
	for _, h := range hc.history {
		if h.BookID == b.ID {
			return true
		}
	}
	return false
}
