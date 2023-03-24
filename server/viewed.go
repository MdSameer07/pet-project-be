package server

import (
	"context"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) MarkAsRead(ctx context.Context, req *proto.MarkAsReadRequest) (*proto.MarkAsReadResponse, error) {
	if req.UserId == 0 || req.MovieId == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Please enter Both UserId and MovieId")
	}

	viewed, err := m.Db.MarkAsRead(req)
	if err != nil {
		return nil, err
	}

	resp := &proto.MarkAsReadResponse{
		Viewed: &proto.Viewed{
			Id:      uint32(viewed.ID),
			UserId:  uint32(viewed.User_Id),
			MovieId: uint32(viewed.Movie_Id),
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) MarkAsUnread(ctx context.Context, req *proto.MarkAsUnreadRequest) (*proto.MarkAsUnreadResponse, error) {
	if req.UserId==0 || req.MovieId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter both UserId and MovieId")
	}

	status, err := m.Db.MarkAsUnread(req)
	if err != nil {
		resp := &proto.MarkAsUnreadResponse{
			Status: status,
			Errors: err.Error(),
		}
		return resp, err
	}

	resp := &proto.MarkAsUnreadResponse{
		Status: status,
		Errors: "",
	}

	return resp, nil
}
