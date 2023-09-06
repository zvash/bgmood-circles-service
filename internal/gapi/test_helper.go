package gapi

import (
	"github.com/jaswdr/faker"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func randomHeaderUser() HeaderUser {
	f := faker.New()
	avatar := f.Internet().URL()
	return HeaderUser{
		Id:         f.UUID().V4(),
		Email:      f.Internet().Email(),
		Name:       f.Person().Name(),
		Password:   "",
		Avatar:     &avatar,
		IsVerified: true,
		CreatedAt:  timestamppb.New(time.Now()),
	}
}
