package gapi

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/util"
	"github.com/zvash/bgmood-circles-service/internal/worker"
	"google.golang.org/grpc/metadata"
	"testing"
)

func newTestServer(t *testing.T, store db.DataStore, messagePublisher worker.MessagePublisher) *Server {
	config, err := util.LoadConfig("../..")
	require.NoError(t, err)
	server, err := NewServer(config, store, messagePublisher)
	require.NoError(t, err)

	return server
}

func newContextWithHeaderUser(t *testing.T, headerUser HeaderUser) context.Context {
	header, err := json.Marshal(&headerUser)
	require.NoError(t, err)
	headerAsString := string(header)
	headers := make(map[string]string)
	headers[userObjectHeader] = headerAsString
	md := metadata.New(headers)
	return metadata.NewIncomingContext(context.Background(), md)
}
