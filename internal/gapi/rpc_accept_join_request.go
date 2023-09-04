package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"log"
)

func (server *Server) AcceptJoinRequest(ctx context.Context, req *cpb.AcceptJoinRequestRequest) (*cpb.AcceptJoinRequestResponse, error) {
	dto := cpbAcceptJoinRequestRequestToValAcceptJoinRequestRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleJoinReq, err := server.db.GetJoinRequest(ctx, dto.JoinRequestID)
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	_, err = server.getCircleForOwner(ctx, circleJoinReq.CircleID)
	if err != nil {
		return nil, err
	}
	_, err = server.db.AcceptJoinRequestTransaction(ctx, db.AcceptJoinRequestTransactionParams{
		MemberID: circleJoinReq.UserID,
		CircleID: circleJoinReq.CircleID,
		AfterAccept: func(member repository.CircleMember) error {
			//TODO: send notification to the new member
			log.Printf("send notification to the new member: %v", member)
			return nil
		},
	})
	if err != nil {
		return nil, returnDBError(err)
	}
	resp := &cpb.AcceptJoinRequestResponse{
		Message: "join request is accepted",
	}
	return resp, nil
}
