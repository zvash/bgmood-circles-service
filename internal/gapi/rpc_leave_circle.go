package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LeaveCircle(ctx context.Context, req *cpb.LeaveCircleRequest) (*cpb.LeaveCircleResponse, error) {
	dto := cpbLeaveCircleRequestToValLeaveCircleRequest(req)
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
	circle, err := server.db.GetCircle(ctx, circleUUID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	if circle.OwnerID == userUUID {
		return nil, status.Errorf(codes.FailedPrecondition, "owner of circle cannot leave the circle")
	}
	err = server.db.LeaveCircle(ctx, repository.LeaveCircleParams{
		CircleID: circle.ID,
		MemberID: userUUID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	return &cpb.LeaveCircleResponse{
		Messsage: "you successfully left the circle.",
	}, nil
}
