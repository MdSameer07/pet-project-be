package server

import (
	"context"
	
	"example.com/pet-project/proto"
)

func (m *MoviesuggestionsServiceserver) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {

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