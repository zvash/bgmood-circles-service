package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) GetCircleJoinRequests(ctx context.Context, req *cpb.GetCircleJoinRequestsRequest) (*cpb.GetCircleJoinRequestsResponse, error) {
	dto := cpbGetCircleJoinRequestsRequestToValGetCircleJoinRequestsRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	_, err = server.getCircleForOwner(ctx, circleUUID)
	if err != nil {
		return nil, err
	}
	var limit int32 = 10
	joinRequests, err := server.db.GetCircleJoinRequests(ctx, repository.GetCircleJoinRequestsParams{
		CircleID: circleUUID,
		Offset:   int32(dto.Page-1) * limit,
		Limit:    limit,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	userIds := make([]string, 0)
	for _, record := range joinRequests {
		userIds = append(userIds, record.UserID.String())
	}
	resp := &cpb.GetCircleJoinRequestsResponse{
		UserIds: userIds,
	}
	return resp, nil
}
