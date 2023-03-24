package server

import (
	"context"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) AddReviewForMovie(ctx context.Context, req *proto.AddReviewRequest) (*proto.AddReviewResponse, error) {
	if req.UserId==0 || req.MovieId==0 || req.Description==""{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter UserId, MovieId and Description")
	}

	if req.Stars > 5 || req.Stars<=0{
		return nil, status.Errorf(codes.FailedPrecondition, "Stars should be given in the range of 1 to 5")
	}

	review,err := m.Db.AddReviewForMovie(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.AddReviewResponse{
		Review: &proto.Review{
			Id:          uint32(review.ID),
			UserId:      uint32(review.User_Id),
			MovieId:     uint32(review.Movie_Id),
			Description: review.Description,
			Stars:       uint32(review.Stars),
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) UpdateReviewForMovie(ctx context.Context, req *proto.UpdateReviewRequest) (*proto.UpdateReviewResponse, error) {

	if req.UserId==0 || req.MovieId==0 || req.Description==""{
		return nil,status.Errorf(codes.FailedPrecondition,"Enter userId movieId and Description to update")
	}

	review,err := m.Db.UpdateReviewForMovie(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.UpdateReviewResponse{
		Review: &proto.Review{
			UserId:      uint32(review.User_Id),
			MovieId:     uint32(review.Movie_Id),
			Description: review.Description,
			Stars:       uint32(review.Stars),
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) DeleteReviewForMovie(ctx context.Context, req *proto.DeleteReviewRequest) (*proto.DeleteReviewResponse, error) {
	if req.UserId==0 || req.MovieId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter valid UserId and MovieId")
	}

	status,err := m.Db.DeleteReviewForMovie(req)
	if err!=nil{
		resp := &proto.DeleteReviewResponse{
			Status: status,
			Errors: err.Error(),
		}
		return resp,err
	}

	resp := &proto.DeleteReviewResponse{
		Status: status,
		Errors: "",
	}

	return resp, nil
}