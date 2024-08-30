package recommendation

import (
	"context"
	"errors"
	"slices"

	"github.com/hrvadl/book-service/internal/domain/book"
	"github.com/hrvadl/book-service/internal/domain/history"
	"github.com/hrvadl/book-service/internal/domain/review"
	"github.com/hrvadl/book-service/internal/domain/user"
)

const (
	RatingOK        = 3
	RatingGood      = 4
	RatingExcellent = 5
)

func NewService(rs ReviewSource, rhs ReadingHistorySource, us UserSource, bs BookSource) *Service {
	return &Service{
		reviews: rs,
		history: rhs,
		users:   us,
		books:   bs,
	}
}

//go:generate mockgen -destination=./mocks/mock_reviews.go -package=mocks . ReviewSource
type ReviewSource interface {
	GetByUserID(ctx context.Context, userID int) ([]review.Review, error)
}

//go:generate mockgen -destination=./mocks/mock_history.go -package=mocks . ReadingHistorySource
type ReadingHistorySource interface {
	GetByUserID(ctx context.Context, userID int) ([]history.ReadingHistory, error)
}

//go:generate mockgen -destination=./mocks/mock_users.go -package=mocks . UserSource
type UserSource interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
}

//go:generate mockgen -destination=./mocks/mock_books.go -package=mocks . BookSource
type BookSource interface {
	GetAll(ctx context.Context) ([]book.Book, error)
}

type Service struct {
	reviews ReviewSource
	history ReadingHistorySource
	users   UserSource
	books   BookSource
}

func (s *Service) GetRecommendedBookFor(ctx context.Context, userID int) (*book.Book, error) {
	reviews, err := s.reviews.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	history, err := s.history.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	user, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	books, err := s.books.GetAll(ctx)
	if err != nil {
		return nil, errors.Join(ErrFailedToGet, err)
	}

	// TODO: use specification?
	unreadBooks := s.getUnreadBooks(history, books)
	preferredGenres := user.PreferredGenres
	likedAuthorsIDs := s.getLikedAuthors(reviews)

	recommendedBook := s.getBookWithPrefferedGenresOrAuthors(
		unreadBooks,
		preferredGenres,
		likedAuthorsIDs,
	)
	if recommendedBook == nil {
		return nil, ErrNoRecommendation
	}

	return recommendedBook, nil
}

func (s *Service) getBookWithPrefferedGenresOrAuthors(
	books []book.Book,
	genres []string,
	authorIDs []int,
) *book.Book {
	for _, b := range books {
		for _, g := range genres {
			if slices.Contains(b.Genres, g) {
				return &b
			}
		}

		for _, a := range b.Authors {
			if slices.Contains(authorIDs, a.ID) {
				return &b
			}
		}
	}

	return nil
}

func (s *Service) getUnreadBooks(history []history.ReadingHistory, books []book.Book) []book.Book {
	unreadBooks := make([]book.Book, 0, len(books)/2)
	historyComparator := newHistoryComparator(history)
	for _, b := range books {
		if !historyComparator.Contains(b) {
			unreadBooks = append(unreadBooks, b)
		}
	}
	return unreadBooks
}

func (s *Service) getLikedAuthors(reviews []review.Review) []int {
	likedAuthors := make(map[int]struct{})
	for _, r := range reviews {
		if r.Rating > RatingOK {
			likedAuthors[r.AuthorID] = struct{}{}
		}
	}

	likedAuthorsRes := make([]int, 0, len(likedAuthors))
	for a := range likedAuthors {
		likedAuthorsRes = append(likedAuthorsRes, a)
	}

	return likedAuthorsRes
}
