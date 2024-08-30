package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/user"
	"github.com/hrvadl/book-service/internal/domain/user/mocks"
)

func TestServiceGetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}

	type fields struct {
		users func(*gomock.Controller) user.UserSource
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		want    *user.User
		wantErr bool
	}{
		{
			name: "Should return user correctly",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			fields: fields{
				users: func(c *gomock.Controller) user.UserSource {
					us := mocks.NewMockUserSource(c)
					us.EXPECT().GetByID(context.Background(), 1).Times(1).Return(&user.User{
						ID:              1,
						Name:            "Vadym",
						Email:           "vadym@vadym.com",
						PreferredGenres: []string{"horror", "tech"},
					}, nil)
					return us
				},
			},
			wantErr: false,
			want: &user.User{
				ID:              1,
				Name:            "Vadym",
				Email:           "vadym@vadym.com",
				PreferredGenres: []string{"horror", "tech"},
			},
		},
		{
			name: "Should return domain error when failed",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			fields: fields{
				users: func(c *gomock.Controller) user.UserSource {
					us := mocks.NewMockUserSource(c)
					us.EXPECT().
						GetByID(context.Background(), 1).
						Times(1).
						Return(nil, errors.New("failed"))
					return us
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := user.NewService(tt.fields.users(gomock.NewController(t)))
			got, err := s.GetByID(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				require.ErrorIs(t, err, user.ErrFailedToGet)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
