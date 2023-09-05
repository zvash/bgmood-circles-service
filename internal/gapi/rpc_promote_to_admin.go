package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) PromoteToAdmin(
	ctx context.Context,
	req *cpb.PromoteToAdminRequest,
) (*cpb.AccessModificationResponse, error) {
	dto := cpbPromoteToAdminRequestToValPromoteToAdminRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, memberUUID, err := server.prepareChangingAccessToCircle(ctx, dto.CircleID, dto.UserID)
	if err != nil {
		return nil, err
	}
	circleMember, err := server.db.SetCircleAccessToAdmin(ctx, repository.SetCircleAccessToAdminParams{
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
