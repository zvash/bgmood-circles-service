package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) PromoteToPoster(
	ctx context.Context,
	req *cpb.PromoteToPosterRequest,
) (*cpb.AccessModificationResponse, error) {
	dto := cpbPromoteToPosterRequestToValPromoteToPosterRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, memberUUID, err := server.prepareChangingAccessToCircle(ctx, dto.CircleID, dto.UserID)
	if err != nil {
		return nil, err
	}
	circleMember, err := server.db.SetCircleAccessToPoster(ctx, repository.SetCircleAccessToPosterParams{
		CircleID: circleUUID,
		MemberID: memberUUID,
	})
	if err != nil {
		return nil, returnDBError(err)
	}
	resp := &cpb.AccessModificationResponse{
		CircleId:      circleMember.CircleID.String(),
		CurrentAccess: string(circleMember.AcceptanceType),
	}
	return resp, nil
}
