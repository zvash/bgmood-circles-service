package gapi

import (
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/util"
	"github.com/zvash/bgmood-circles-service/internal/val"
	"github.com/zvash/bgmood-circles-service/internal/worker"
)

// Server serves gRPC requests for our auth service.
type Server struct {
	cpb.UnimplementedCirclesServer
	config           util.Config
	db               db.DataStore
	validator        *val.XValidator
	messagePublisher worker.MessagePublisher
}

func NewServer(config util.Config, dataStore db.DataStore, messagePublisher worker.MessagePublisher) (*Server, error) {
	server := &Server{
		config:           config,
		db:               dataStore,
		validator:        val.NewValidator(),
		messagePublisher: messagePublisher,
	}

	return server, nil
}
