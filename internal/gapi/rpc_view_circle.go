package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
)

func (server *Server) ViewCircle(ctx context.Context, req *cpb.ViewCircleRequest) (*cpb.ViewCircleResponse, error) {
	dto := cpbViewCircleRequestToValViewCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	circleMeta, err := server.getCircleMeta(ctx, circleUUID, userUUID)
	if err != nil {
		return nil, err
	}

	resp := &cpb.ViewCircleResponse{
		Circle: circleMeta,
	}
	return resp, nil
}
