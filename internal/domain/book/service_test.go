package book_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/author"
	"github.com/hrvadl/book-service/internal/domain/book"
	"github.com/hrvadl/book-service/internal/domain/book/mocks"
)

func TestAddBook(t *testing.T) {
	t.Parallel()
	type fields struct {
		books func(*gomock.Controller) book.BookSource
	}

	type args struct {
		ctx  context.Context
		book book.Book
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "Should save book correctly",
			args: args{
				ctx: context.Background(),
				book: book.Book{
					Authors: []author.Author{
						{
							Name:    "Vadym",
							Surname: "Hrashchenko",
							BirthCountry: author.Country{
								Name: "Ukraine",
							},
						},
					},
				},
			},
			fields: fields{
				books: func(c *gomock.Controller) book.BookSource {
					book := book.Book{
						Authors: []author.Author{
							{
								Name:    "Vadym",
								Surname: "Hrashchenko",
								BirthCountry: author.Country{
									Name: "Ukraine",
								},
							},
						},
					}
					s := mocks.NewMockBookSource(c)
					s.EXPECT().Save(gomock.Any(), book).Times(1).Return(1, nil)
					return s
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Should return domain error if failed",
			args: args{
				ctx: context.Background(),
				book: book.Book{
					Authors: []author.Author{
						{
							Name:    "Vadym",
							Surname: "Hrashchenko",
							BirthCountry: author.Country{
								Name: "Ukraine",
							},
						},
					},
				},
			},
			fields: fields{
				books: func(c *gomock.Controller) book.BookSource {
					book := book.Book{
						Authors: []author.Author{
							{
								Name:    "Vadym",
								Surname: "Hrashchenko",
								BirthCountry: author.Country{
									Name: "Ukraine",
								},
							},
						},
					}
					s := mocks.NewMockBookSource(c)
					s.EXPECT().Save(gomock.Any(), book).Times(1).Return(0, errors.New("failed"))
					return s
				},
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := book.NewService(tt.fields.books(gomock.NewController(t)))
			id, err := svc.Add(tt.args.ctx, tt.args.book)
			if tt.wantErr {
				require.ErrorIs(t, err, book.ErrFailedToAdd)
				return
			}

			require.Equal(t, tt.want, id)
		})
	}
}

func TestServiceGetAll(t *testing.T) {
	t.Parallel()
	type fields struct {
		books func(*gomock.Controller) book.BookSource
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []book.Book
		wantErr bool
	}{
		{
			name: "Should return books",
			fields: fields{
				books: func(c *gomock.Controller) book.BookSource {
					s := mocks.NewMockBookSource(c)
					s.EXPECT().GetAll(context.Background()).Times(1).Return([]book.Book{
						{
							ID:    1,
							Title: "Test",
						},
						{
							ID:    2,
							Title: "Test 2",
						},
					}, nil)
					return s
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: []book.Book{
				{
					ID:    1,
					Title: "Test",
				},
				{
					ID:    2,
					Title: "Test 2",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return domain error when source failed",
			fields: fields{
				books: func(c *gomock.Controller) book.BookSource {
					s := mocks.NewMockBookSource(c)
					s.EXPECT().
						GetAll(context.Background()).
						Times(1).
						Return(nil, errors.New("failed"))
					return s
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := book.NewService(tt.fields.books(gomock.NewController(t)))
			got, err := s.GetAll(tt.args.ctx)

			if tt.wantErr {
				require.ErrorIs(t, err, book.ErrFailedToGet)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
