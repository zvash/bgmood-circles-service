package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) SetCircleWPChangeAccess(ctx context.Context,
	req *cpb.SetCircleWPChangeAccessRequest,
) (*cpb.CircleWPChangeAccessResponse, error) {
	dto := cpbSetCircleWPChangeAccessRequestToValSetCircleWPChangeAccessRequest(req)
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
	_, err = server.db.GetCircleMember(ctx, repository.GetCircleMemberParams{
		MemberID: userUUID,
		CircleID: circleUUID,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	dbAccessType := repository.AcceptanceTypeASKFIRST
	switch dto.AccessType {
	case "ASK_FIRST":
		dbAccessType = repository.AcceptanceTypeASKFIRST
	case "ACCEPT_ALL":
		dbAccessType = repository.AcceptanceTypeACCEPTALL
	case "REJECT_ALL":
		dbAccessType = repository.AcceptanceTypeREJECTALL
	}
	membership, err := server.db.SetCircleMemberAccess(ctx, repository.SetCircleMemberAccessParams{
		CircleID:       circleUUID,
		MemberID:       userUUID,
		AcceptanceType: dbAccessType,
	})
	if err != nil {
		return nil, returnDBError(err)
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
