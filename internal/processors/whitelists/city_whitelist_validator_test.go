package whitelists

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"rebrainme/gotest/internal/entities"
	"rebrainme/gotest/test/mocks/packages/clientsrepomocks"
)

func TestCityWhitelistValidator_Validate(t *testing.T) {
	type mocks struct {
		clientRepoMocks *clientsrepomocks.MockRepository
	}

	req := require.New(t)

	clientEmail := "example@example.com"

	cases := map[string]struct {
		setup         func(ctx context.Context, m mocks)
		grantedCities []string
		err           error
	}{
		"success validate": {
			setup: func(ctx context.Context, m mocks) {
				m.clientRepoMocks.EXPECT().
					FindClientByEmail(ctx, clientEmail).
					Return(&entities.Client{City: "Moscow"}, nil).
					Times(1)
			},
			grantedCities: []string{"Moscow"},
			err:           nil,
		},
		"fail validate with not granted city": {
			setup: func(ctx context.Context, m mocks) {
				m.clientRepoMocks.EXPECT().
					FindClientByEmail(ctx, clientEmail).
					Return(&entities.Client{City: "NY"}, nil).
					Times(1)
			},
			grantedCities: []string{"Moscow"},
			err:           errors.New("client city is not granted"),
		},
		"fail validate with find client error": {
			setup: func(ctx context.Context, m mocks) {
				m.clientRepoMocks.EXPECT().
					FindClientByEmail(ctx, clientEmail).
					Return(nil, sql.ErrNoRows).
					Times(1)
			},
			grantedCities: []string{"Moscow"},
			err:           sql.ErrNoRows,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockContainer := mocks{
				clientRepoMocks: clientsrepomocks.NewMockRepository(mockCtrl),
			}
			cs.setup(ctx, mockContainer)

			validator := newCityWhitelistValidator(mockContainer.clientRepoMocks, cs.grantedCities)
			err := validator.Validate(ctx, clientEmail)
			req.Equal(cs.err, err)
		})
	}
}
