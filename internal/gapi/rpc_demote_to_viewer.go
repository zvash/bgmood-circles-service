package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) DemoteToViewer(
	ctx context.Context,
	req *cpb.DemoteToViewerRequest,
) (*cpb.AccessModificationResponse, error) {
	dto := cpbDemoteToViewerRequestToValDemoteToViewerRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, memberUUID, err := server.prepareChangingAccessToCircle(ctx, dto.CircleID, dto.UserID)
	if err != nil {
		return nil, err
	}
	circleMember, err := server.db.SetCircleAccessToViewer(ctx, repository.SetCircleAccessToViewerParams{
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
