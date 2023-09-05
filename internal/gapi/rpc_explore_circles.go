package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
)

func (server *Server) ExploreCircles(ctx context.Context, req *cpb.ExploreCirclesRequest) (*cpb.ExploreCirclesResponse, error) {
	dto := cpbExploreCirclesRequestToValExploreCirclesRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}
	circlesMeta, err := server.getExploredCirclesMeta(ctx, userUUID, dto.Page)
	if err != nil {
		return nil, err
	}
	return &cpb.ExploreCirclesResponse{
		Circles: circlesMeta,
	}, nil
}
