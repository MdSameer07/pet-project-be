package server

import (
	"context"
	"log"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	if req.Id==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Enter id of user to be updated")
	}

	log.Print(req.Id)

	user,err := m.Db.UpdateProfile(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.UpdateProfileResponse{
		Id:          uint32(user.ID),
		Name:        user.UserName,
		Email:       user.UserEmail,
		PhoneNumber: user.UserPhoneNumber,
	}

	return resp, nil
}