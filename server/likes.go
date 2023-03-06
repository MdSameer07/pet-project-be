package server
import (
	"context"

	"example.com/pet-project/proto"
)


func (m *MoviesuggestionsServiceserver) AddMovieToLikes(ctx context.Context, req *proto.AddMovieToLikesRequest) (*proto.AddMovieToLikesResponse, error) {

	likes,err := m.Db.AddMovieToLikes(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.AddMovieToLikesResponse{
		Like: &proto.Likes{
			Id:      uint32(likes.ID),
			UserId:  uint32(likes.User_Id),
			MovieId: uint32(likes.Movie_Id),
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) RemoveMovieFromLikes(ctx context.Context, req *proto.RemoveMovieFromLikesRequest) (*proto.RemoveMovieFromLikesResponse, error) {

	status,err := m.Db.RemoveMovieFromLikes(req)
	if err!=nil{
		resp := &proto.RemoveMovieFromLikesResponse{
			Status: status,
			Errors: err.Error(),
		}
		return resp,err
	}

	resp := &proto.RemoveMovieFromLikesResponse{
		Status: status,
		Errors: "",
	}

	return resp, nil
}