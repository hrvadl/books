package history_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/history"
	"github.com/hrvadl/book-service/internal/domain/history/mocks"
)

func TestHistoryGetByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
	}

	type fields struct {
		history func(*gomock.Controller) history.HistorySource
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
		want    []history.ReadingHistory
	}{
		{
			name: "Should return history by userID",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			fields: fields{
				history: func(c *gomock.Controller) history.HistorySource {
					s := mocks.NewMockHistorySource(c)
					s.EXPECT().
						GetByUserID(context.Background(), 1).
						Times(1).
						Return([]history.ReadingHistory{
							{
								BookID: 1,
								UserID: 1,
							},
							{
								BookID: 2,
								UserID: 1,
							},
						}, nil)
					return s
				},
			},
			wantErr: false,
			want: []history.ReadingHistory{
				{
					BookID: 1,
					UserID: 1,
				},
				{
					BookID: 2,
					UserID: 1,
				},
			},
		}, {
			name: "Should return domain error when failed",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			fields: fields{
				history: func(c *gomock.Controller) history.HistorySource {
					s := mocks.NewMockHistorySource(c)
					s.EXPECT().
						GetByUserID(context.Background(), 1).
						Times(1).
						Return(nil, errors.New("failed"))
					return s
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := history.NewService(tt.fields.history(gomock.NewController(t)))
			got, err := s.GetByUserID(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.ErrorIs(t, err, history.ErrFailedToGet)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
