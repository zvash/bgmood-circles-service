package gapi

import (
	"context"
	"errors"
	"github.com/zvash/bgmood-circles-service/internal/cpb"
	"github.com/zvash/bgmood-circles-service/internal/db"
	"github.com/zvash/bgmood-circles-service/internal/db/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (server *Server) InviteUserToCircle(
	ctx context.Context,
	req *cpb.InviteUserToCircleRequest,
) (*cpb.InviteUserToCircleResponse, error) {
	dto := cpbInviteUserToCircleRequestToValInviteUserToCircleRequest(req)
	if errs := server.validator.Validate(dto); errs != nil {
		return nil, errorResponsesToStatusErrors(errs)
	}
	circleUUID, err := stringToUUID(dto.CircleID)
	if err != nil {
		return nil, err
	}
	userUUID, err := stringToUUID(dto.UserID)
	if err != nil {
		return nil, err
	}
	_, err = server.getCircleForOwner(ctx, circleUUID)
	if err != nil {
		return nil, err
	}
	transactionParams := db.InviteToCircleTransactionParams{
		InviteToCircleParams: repository.InviteToCircleParams{
			CircleID: circleUUID,
			UserID:   userUUID,
		},
		AfterInvite: func(invitation repository.CircleInvitation) error {
			//TODO: send the invitation
			log.Printf("Send the invitation")
			return nil
		},
	}
	_, err = server.db.InviteToCircleTransaction(ctx, transactionParams)
	if err != nil {
		if errors.Is(err, db.ErrUniqueViolation) {
			return nil, status.Errorf(codes.AlreadyExists, "user is already invited to this circle")
		}
		return nil, status.Errorf(codes.Internal, "internal server error.")
	}
	resp := &cpb.InviteUserToCircleResponse{
		Message: "An invitation to this circle will be sent to this user",
	}
	return resp, nil
}
