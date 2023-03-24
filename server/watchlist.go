package server

import (
	"context"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) AddMovieToWatchList(ctx context.Context, req *proto.AddMovieToWatchListRequest) (*proto.AddMovieToWatchListResponse, error) {
	if req.UserId==0 || req.MovieId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter UserId and MovieId")
	}

	watchlist,err := m.Db.AddMovieToWatchList(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.AddMovieToWatchListResponse{
		Watchlist: &proto.WatchList{
			Id:      uint32(watchlist.ID),
			UserId:  uint32(watchlist.User_Id),
			MovieId: uint32(watchlist.Movie_Id),
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) RemoveMovieFromWatchList(ctx context.Context, req *proto.RemoveMovieFromWatchListRequest) (*proto.RemoveMovieFromWatchListResponse, error) {
	if req.UserId==0 || req.MovieId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter both UserId and MovieId")
	}

	status,err := m.Db.RemoveMovieFromWatchList(req)
	if err!=nil{
		resp := &proto.RemoveMovieFromWatchListResponse{
			Status: status,
			Errors: err.Error(),
		}
		return resp,err
	}

	resp := &proto.RemoveMovieFromWatchListResponse{
		Status: status,
		Errors: "",
	}

	return resp, nil
}