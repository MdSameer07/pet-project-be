package server

import (
	"context"

	"example.com/pet-project/proto"
)

func (m *MoviesuggestionsServiceserver) GiveFeedBack(ctx context.Context, req *proto.GiveFeedBackRequest) (*proto.GiveFeedBackResponse, error) {

	feedback,err := m.Db.GiveFeedBack(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.GiveFeedBackResponse{
		Feedback: &proto.FeedBack{
			Id:          uint32(feedback.ID),
			UserId:      uint32(feedback.User_Id),
			Description: feedback.Description,
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) UpdateFeedBack(ctx context.Context, req *proto.UpdateFeedBackRequest) (*proto.UpdateFeedBackResponse, error) {

	feedback,err := m.Db.UpdateFeedBack(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.UpdateFeedBackResponse{
		Feedback: &proto.FeedBack{
			Id:          uint32(feedback.ID),
			UserId:      uint32(feedback.User_Id),
			Description: feedback.Description,
		},
	}

	return resp, nil
}

func (m *MoviesuggestionsServiceserver) DeleteFeedBack(ctx context.Context, req *proto.DeleteFeedBackRequest) (*proto.DeleteFeedBackResponse, error) {

	status,err := m.Db.DeleteFeedBack(req)
	if err!=nil{
		resp := &proto.DeleteFeedBackResponse{
			Status: status,
			Errors: err.Error(),
		}
		return resp,err
	}

	resp := &proto.DeleteFeedBackResponse{
		Status: status,
		Errors: "",
	}

	return resp, nil
}