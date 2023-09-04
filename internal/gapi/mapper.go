package gapi

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"github.com/zvash/bgmood-circles-service/internal/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorResponseToErrorDetailsBadRequestFieldViolation(er val.ErrorResponse) *errdetails.BadRequest_FieldViolation {
	fieldName := strcase.ToSnake(er.FailedField)
	return &errdetails.BadRequest_FieldViolation{
		Field: fieldName,
		Description: fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			fieldName,
			er.Value,
			er.Tag,
		),
	}
}

func errorResponsesToErrorDetailsBadRequestFieldViolations(ers []val.ErrorResponse) (violations []*errdetails.BadRequest_FieldViolation) {
	for _, er := range ers {
		violations = append(violations, errorResponseToErrorDetailsBadRequestFieldViolation(er))
	}
	return
}

func errorResponsesToStatusErrors(errs []val.ErrorResponse) error {
	violations := errorResponsesToErrorDetailsBadRequestFieldViolations(errs)
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}
	return statusDetails.Err()
}

func cpbCreateCircleRequestToValCreateCircleRequest(req *cpb.CreateCircleRequest) val.CreateCircleRequest {
	return val.CreateCircleRequest{
		Title:       req.Title,
		Avatar:      req.Avatar,
		Description: req.GetDescription(),
		CircleType:  req.GetCircleType(),
		IsPrivate:   req.GetIsPrivate(),
	}
}

func dbCircleToCPBCircle(circle repository.Circle) *cpb.Circle {
	cCircle := &cpb.Circle{
		Id:         circle.ID.String(),
		OwnerId:    circle.OwnerID.String(),
		Title:      circle.Title,
		Avatar:     circle.Avatar.String,
		CircleType: string(circle.CircleType),
		IsPrivate:  circle.IsPrivate,
		IsFeatured: circle.IsFeatured,
	}
	if circle.Description.Valid {
		cCircle.Description = circle.Description.String
	}
	return cCircle
}

func cpbEditCircleRequestToValEditCircleRequest(req *cpb.EditCircleRequest) val.EditCircleRequest {
	return val.EditCircleRequest{
		Title:       req.GetTitle(),
		Avatar:      req.GetAvatar(),
		Description: req.GetDescription(),
		CircleType:  req.GetCircleType(),
		IsPrivate:   req.GetIsPrivate(),
	}
}

func cpbInviteUserToCircleRequestToValInviteUserToCircleRequest(req *cpb.InviteUserToCircleRequest) val.InviteUserToCircleRequest {
	return val.InviteUserToCircleRequest{
		CircleID: req.GetCircleId(),
		UserID:   req.GetUserId(),
	}
}

func cpbKickFromCircleRequestToValKickFromCircleRequest(req *cpb.KickFromCircleRequest) val.KickFromCircleRequest {
	return val.KickFromCircleRequest{
		CircleID: req.GetCircleId(),
		UserID:   req.GetUserId(),
	}
}

func cpbSendJoinRequestRequestToValSendJoinRequestRequest(req *cpb.SendJoinRequestRequest) val.SendJoinRequestRequest {
	return val.SendJoinRequestRequest{
		CircleID: req.GetCircleId(),
	}
}

func cpbAcceptJoinRequestRequestToValAcceptJoinRequestRequest(req *cpb.AcceptJoinRequestRequest) val.AcceptJoinRequestRequest {
	return val.AcceptJoinRequestRequest{
		JoinRequestID: req.GetJoinRequestId(),
	}
}

func cpbRemoveJoinRequestRequestToValRemoveJoinRequestRequest(req *cpb.RemoveJoinRequestRequest) val.RemoveJoinRequestRequest {
	return val.RemoveJoinRequestRequest{
		JoinRequestID: req.GetJoinRequestId(),
	}
}

func cpbLeaveCircleRequestToValLeaveCircleRequest(req *cpb.LeaveCircleRequest) val.LeaveCircleRequest {
	return val.LeaveCircleRequest{
		CircleID: req.GetCircleId(),
	}
}
