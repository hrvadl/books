package author_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/author"
	"github.com/hrvadl/book-service/internal/domain/author/mocks"
)

func TestAddAuthor(t *testing.T) {
	type fields struct {
		authors func(*gomock.Controller) author.AuthorSaver
	}

	type args struct {
		ctx    context.Context
		author author.Author
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "Should add author correctly",
			args: args{
				ctx: context.Background(),
				author: author.Author{
					Name:        "Name",
					Surname:     "Surname",
					DateOfBirth: time.Date(2004, 0o7, 25, 0, 0, 0, 0, time.FixedZone("us", 3)),
				},
			},
			fields: fields{
				authors: func(c *gomock.Controller) author.AuthorSaver {
					s := mocks.NewMockAuthorSaver(c)
					s.EXPECT().Save(gomock.Any(), author.Author{
						Name:        "Name",
						Surname:     "Surname",
						DateOfBirth: time.Date(2004, 0o7, 25, 0, 0, 0, 0, time.FixedZone("us", 3)),
					}).Times(1).Return(1, nil)
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
				author: author.Author{
					Name:        "Name",
					Surname:     "Surname",
					DateOfBirth: time.Date(2004, 0o7, 25, 0, 0, 0, 0, time.FixedZone("us", 3)),
				},
			},
			fields: fields{
				authors: func(c *gomock.Controller) author.AuthorSaver {
					s := mocks.NewMockAuthorSaver(c)
					s.EXPECT().Save(gomock.Any(), author.Author{
						Name:        "Name",
						Surname:     "Surname",
						DateOfBirth: time.Date(2004, 0o7, 25, 0, 0, 0, 0, time.FixedZone("us", 3)),
					}).Times(1).Return(0, errors.New("failed"))
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
			svc := author.NewService(tt.fields.authors(gomock.NewController(t)))
			id, err := svc.Add(tt.args.ctx, tt.args.author)
			if tt.wantErr {
				require.ErrorIs(t, err, author.ErrFailedToAdd)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, id)
		})
	}
}
