package gapi

import (
	"context"
	"errors"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SendJoinRequest(ctx context.Context, req *cpb.SendJoinRequestRequest) (*cpb.SendJoinRequestResponse, error) {
	dto := cpbSendJoinRequestRequestToValSendJoinRequestRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized.")
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	circle, err := server.db.GetCircle(ctx, circleUUID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "circle was not found.")
		}
		return nil, internalServerError()
	}
	if circle.IsPrivate && circle.CircleType == repository.CircleTypeHALL {
		_, err = server.db.CreateJoinRequest(ctx, repository.CreateJoinRequestParams{
			CircleID: circleUUID,
			UserID:   userUUID,
		})
		if err != nil {
			return nil, uniqueOrInternalError(err)
		}
		resp := &cpb.SendJoinRequestResponse{
			Message: "Join request was made.",
		}
		return resp, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "you cannot send join request to this circle")

}
