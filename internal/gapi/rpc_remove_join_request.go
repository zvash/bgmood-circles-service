package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) RemoveJoinRequest(
	ctx context.Context,
	req *cpb.RemoveJoinRequestRequest,
) (*cpb.RemoveJoinRequestResponse, error) {
	dto := cpbRemoveJoinRequestRequestToValRemoveJoinRequestRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}
	joinReq, err := server.db.GetJoinRequest(ctx, dto.JoinRequestID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	//take back your own join request
	if userUUID == joinReq.UserID {
		return server.deleteJoinRequestRecord(ctx, joinReq)
	}
	circle, err := server.db.GetCircle(ctx, joinReq.CircleID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	//reject join request by circle owner
	_, err = server.getCircleForOwner(ctx, circle.ID)
	if err != nil {
		return nil, err
	}
	return server.deleteJoinRequestRecord(ctx, joinReq)
}

func (server *Server) deleteJoinRequestRecord(ctx context.Context, joinReq repository.CircleJoinRequest) (*cpb.RemoveJoinRequestResponse, error) {
	err := server.db.RemoveJoinRequest(ctx, repository.RemoveJoinRequestParams{
		UserID:   joinReq.UserID,
		CircleID: joinReq.CircleID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	return &cpb.RemoveJoinRequestResponse{
		Message: "Join request to the circle is removed.",
	}, nil
}
