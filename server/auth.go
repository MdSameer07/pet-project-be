package server

import (
	"context"

	"example.com/pet-project/proto"
)

func (m *MoviesuggestionsServiceserver) Login(ctx context.Context,req *proto.LoginRequest) (*proto.LoginResponse,error){

	Id,err := m.Db.Login(req)

	if err!=nil{
		resp := &proto.LoginResponse{
			Id:0,
		}
		return resp,err
	}

	resp := &proto.LoginResponse{
		Id: Id,
	}
	return resp,err
}

func (m *MoviesuggestionsServiceserver) Register(ctx context.Context,req *proto.RegisterRequest) (*proto.RegisterResponse,error){

	user,err := m.Db.Register(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.RegisterResponse{
		User: &proto.User{
			Id: uint32(user.ID),
			Name: user.UserName,
			Email: user.UserEmail,
			PhoneNumber: user.UserPhoneNumber,
		},
	}

	return resp,nil
}