package gapi

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	mockdb "github.com/zvash/bgmood-circles-service/internal/db/mock"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetAvailableReactionsRPC(t *testing.T) {
	user := randomHeaderUser()

	testCases := []struct {
		name          string
		req           *cpb.GetAvailableReactionsRequest
		buildStubs    func(store *mockdb.MockDataStore)
		buildContext  func(t *testing.T, user HeaderUser) context.Context
		checkResponse func(t *testing.T, res *cpb.GetAvailableReactionsResponse, err error)
	}{
		{
			name: "OK",
			req:  &cpb.GetAvailableReactionsRequest{},
			buildStubs: func(store *mockdb.MockDataStore) {
				availableReactions := []repository.Reaction{
					{
						ID:                 1,
						Name:               "Thumbs Up",
						Emoji:              pgtype.Text{String: "👍", Valid: true},
						TextRepresentation: pgtype.Text{String: ":thumbsup:", Valid: true},
					},
					{
						ID:                 2,
						Name:               "Thumbs Down",
						Emoji:              pgtype.Text{String: "👎", Valid: true},
						TextRepresentation: pgtype.Text{String: ":thumbdown:", Valid: true},
					},
				}
				store.EXPECT().GetAvailableReactions(gomock.Any()).
					Times(1).
					Return(availableReactions, nil)
			},
			buildContext: func(t *testing.T, user HeaderUser) context.Context {
				return newContextWithHeaderUser(t, user)
			},
			checkResponse: func(t *testing.T, res *cpb.GetAvailableReactionsResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, 2, len(res.GetReactions()))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockDataStore(storeCtrl)

			tc.buildStubs(store)
			server := newTestServer(t, store, nil)

			ctx := tc.buildContext(t, user)
			res, err := server.GetAvailableReactions(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
