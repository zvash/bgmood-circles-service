package gapi

import (
	"context"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
)

func (server *Server) GetJoinedCircles(
	ctx context.Context,
	req *cpb.GetJoinedCirclesRequest,
) (*cpb.GetJoinedCirclesResponse, error) {
	dto := cpbGetJoinedCirclesRequestToValGetJoinedCirclesRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	userUUID, err := server.getOwnerUUID(ctx)
	if err != nil {
		return nil, err
	}
	var limit int32 = 10
	dbCircles, err := server.db.ListJoinedCirclesPaginated(ctx, repository.ListJoinedCirclesPaginatedParams{
		MemberID: userUUID,
		Offset:   int32(dto.Page-1) * limit,
		Limit:    limit,
	})
	if err != nil {
		return nil, notFoundOrInternalError(err)
	}
	circles := make([]*cpb.Circle, 0)
	for _, dbCircle := range dbCircles {
		circles = append(circles, dbCircleToCPBCircle(dbCircle))
	}
	return &cpb.GetJoinedCirclesResponse{
		Circles: circles,
	}, nil
}
