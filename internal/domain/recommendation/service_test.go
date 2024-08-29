package recommendation_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/author"
	"github.com/hrvadl/book-service/internal/domain/book"
	"github.com/hrvadl/book-service/internal/domain/history"
	"github.com/hrvadl/book-service/internal/domain/recommendation"
	"github.com/hrvadl/book-service/internal/domain/recommendation/mocks"
	"github.com/hrvadl/book-service/internal/domain/review"
	"github.com/hrvadl/book-service/internal/domain/user"
)

var books = []book.Book{
	{
		ID: 1,
		Authors: []author.Author{
			{
				ID:      1,
				Name:    "Taras",
				Surname: "Shevchenko",
				BirthCountry: author.Country{
					Name: "Ukraine",
				},
			},
		},
		Title:  "Super book",
		Genres: []string{"business"},
	},
	{
		ID: 2,
		Authors: []author.Author{
			{
				ID:      2,
				Name:    "Alex",
				Surname: "Shevchenko",
				BirthCountry: author.Country{
					Name: "Ukraine",
				},
			},
		},
		Title:  "Super book 2",
		Genres: []string{"tech"},
	},
	{
		ID: 3,
		Authors: []author.Author{
			{
				ID:      3,
				Name:    "Eugene",
				Surname: "Shevchenko",
				BirthCountry: author.Country{
					Name: "Ukraine",
				},
			},
		},
		Title:  "Super book 3",
		Genres: []string{"productity"},
	},
	{
		ID: 4,
		Authors: []author.Author{
			{
				ID:      4,
				Name:    "Max",
				Surname: "Shevchenko",
				BirthCountry: author.Country{
					Name: "Ukraine",
				},
			},
		},
		Title:  "Super book 4",
		Genres: []string{"tech"},
	},
	{
		ID: 5,
		Authors: []author.Author{
			{
				ID:      5,
				Name:    "Taras",
				Surname: "Shevchenko",
				BirthCountry: author.Country{
					Name: "Ukraine",
				},
			},
		},
		Title:  "Super book 5",
		Genres: []string{"business"},
	},
}

func TestGetRecommendedBook(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
	}

	type fields struct {
		reviews func(*gomock.Controller) recommendation.ReviewSource
		history func(*gomock.Controller) recommendation.ReadingHistorySource
		users   func(*gomock.Controller) recommendation.UserSource
		books   func(*gomock.Controller) recommendation.BookSource
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		want    *book.Book
		wantErr bool
	}{
		{
			name: "Should recommend book correctly",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			fields: fields{
				reviews: func(c *gomock.Controller) recommendation.ReviewSource {
					s := mocks.NewMockReviewSource(c)
					s.EXPECT().
						GetAllByUserID(context.Background(), 1).
						Times(1).
						Return([]review.Review{
							{
								ID:       1,
								BookID:   2,
								AuthorID: 1,
								Rating:   5,
								Text:     "I Loved this book!",
							},
							{
								ID:       1,
								BookID:   3,
								AuthorID: 2,
								Rating:   3,
								Text:     "Not really great one.",
							},
						}, nil)
					return s
				},
				history: func(c *gomock.Controller) recommendation.ReadingHistorySource {
					s := mocks.NewMockReadingHistorySource(c)
					s.EXPECT().
						GetByUserID(context.Background(), 1).
						Times(1).
						Return([]history.ReadingHistory{
							{
								LastOpened: time.Now(),
								BookID:     2,
								UserID:     1,
							},
							{
								LastOpened: time.Now(),
								BookID:     3,
								UserID:     1,
							},
						}, nil)
					return s
				},
				users: func(c *gomock.Controller) recommendation.UserSource {
					s := mocks.NewMockUserSource(c)
					s.EXPECT().GetByID(context.Background(), 1).Times(1).Return(&user.User{
						ID:              1,
						Name:            "Vadym",
						Email:           "vadym@gmail.com",
						PreferredGenres: []string{"tech", "productivity"},
					}, nil)
					return s
				},
				books: func(c *gomock.Controller) recommendation.BookSource {
					s := mocks.NewMockBookSource(c)
					s.EXPECT().GetAll(context.Background()).Times(1).Return(books, nil)
					return s
				},
			},
			wantErr: false,
			want: &book.Book{
				ID: 1,
				Authors: []author.Author{
					{
						ID:      1,
						Name:    "Taras",
						Surname: "Shevchenko",
						BirthCountry: author.Country{
							Name: "Ukraine",
						},
					},
				},
				Title:  "Super book",
				Genres: []string{"business"},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			svc := recommendation.NewService(
				tt.fields.reviews(ctrl),
				tt.fields.history(ctrl),
				tt.fields.users(ctrl),
				tt.fields.books(ctrl),
			)

			got, err := svc.GetRecommendedBookFor(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.ErrorIs(t, err, recommendation.ErrFailedToGet)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
