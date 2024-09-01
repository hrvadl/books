package review_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/hrvadl/book-service/internal/domain/review"
	"github.com/hrvadl/book-service/internal/domain/review/mocks"
)

func TestGetByUserID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		userID int
	}

	type fields struct {
		reviews func(*gomock.Controller) review.ReviewSource
	}

	tc := []struct {
		name    string
		args    args
		fields  fields
		want    []review.Review
		wantErr bool
	}{
		{
			name: "Should get reviews by userID",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			fields: fields{
				reviews: func(c *gomock.Controller) review.ReviewSource {
					s := mocks.NewMockReviewSource(c)
					s.EXPECT().GetByUserID(context.Background(), 1).Times(1).Return([]review.Review{
						{
							ID:       2,
							BookID:   2,
							AuthorID: 4,
							Rating:   4,
							Text:     "Not bad",
						}, {
							ID:       1,
							BookID:   1,
							AuthorID: 1,
							Rating:   5,
							Text:     "Awesome",
						},
					}, nil)
					return s
				},
			},
			wantErr: false,
			want: []review.Review{
				{
					ID:       2,
					BookID:   2,
					AuthorID: 4,
					Rating:   4,
					Text:     "Not bad",
				}, {
					ID:       1,
					BookID:   1,
					AuthorID: 1,
					Rating:   5,
					Text:     "Awesome",
				},
			},
		},
		{
			name: "Should return domain error",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			fields: fields{
				reviews: func(c *gomock.Controller) review.ReviewSource {
					s := mocks.NewMockReviewSource(c)
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
			s := review.NewService(tt.fields.reviews(gomock.NewController(t)))
			got, err := s.GetByUserID(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.ErrorIs(t, err, review.ErrFailedToGet)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
