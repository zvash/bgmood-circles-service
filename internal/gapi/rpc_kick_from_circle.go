package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) KickFromCircle(
	ctx context.Context,
	req *cpb.KickFromCircleRequest,
) (*cpb.KickFromCircleResponse, error) {
	dto := cpbKickFromCircleRequestToValKickFromCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	userUUID, err := stringToUUID(dto.UserID)
	if err != nil {
		return nil, err
	}
	_, err = server.getCircleForOwner(ctx, circleUUID)
	if err != nil {
		return nil, err
	}
	err = server.db.KickFromCircle(ctx, repository.KickFromCircleParams{
		CircleID: circleUUID,
		MemberID: userUUID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "kicking user from the circle was not successful.")
	}
	resp := &cpb.KickFromCircleResponse{
		Message: "Successfully kicked user from the circle.",
	}
	return resp, nil
}
