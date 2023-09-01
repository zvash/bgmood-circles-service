package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const userObjectHeader = "user-object"

type HeaderUser struct {
	Id         string                 `json:"id"`
	Email      string                 `json:"email"`
	Name       string                 `json:"name"`
	Password   string                 `json:"password"`
	Avatar     *string                `json:"avatar"`
	IsVerified bool                   `json:"is_verified"`
	CreatedAt  *timestamppb.Timestamp `json:"created_at"`
}

func (server *Server) extractUserFromContext(ctx context.Context) (HeaderUser, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if usersAsString := md.Get(userObjectHeader); len(usersAsString) > 0 {
			user := HeaderUser{}
			err := json.Unmarshal([]byte(usersAsString[0]), &user)
			if err != nil {
				return user, err
			}
			return user, nil
		}
	}
	return HeaderUser{}, fmt.Errorf("%s key was not found", userObjectHeader)
}
