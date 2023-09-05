package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) GetCircleWPChangeAccess(
	ctx context.Context,
	req *cpb.GetCircleWPChangeAccessRequest,
) (*cpb.CircleWPChangeAccessResponse, error) {
	dto := cpbGetCircleWPChangeAccessRequestToValGetCircleWPChangeAccessRequest(req)
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
	membership, err := server.db.GetCircleMember(ctx, repository.GetCircleMemberParams{
		CircleID: circleUUID,
		MemberID: userUUID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	accessType := cpb.AccessType_ASK_FIRST
	switch membership.AcceptanceType {
	case repository.AcceptanceTypeACCEPTALL:
		accessType = cpb.AccessType_ACCEPT_ALL
	case repository.AcceptanceTypeASKFIRST:
		accessType = cpb.AccessType_ASK_FIRST
	case repository.AcceptanceTypeREJECTALL:
		accessType = cpb.AccessType_REJECT_ALL
	}
	resp := &cpb.CircleWPChangeAccessResponse{
		CircleId:   circleUUID.String(),
		AccessType: accessType,
	}
	return resp, nil
}
