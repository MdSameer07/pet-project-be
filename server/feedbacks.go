package server

import (
	"context"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) GiveFeedBack(ctx context.Context, req *proto.GiveFeedBackRequest) (*proto.GiveFeedBackResponse, error) {
	if req.UserId==0  || req.Description==""{
		return nil,status.Errorf(codes.FailedPrecondition,"Enter both userId and description of the feedback")
	}

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

	if req.UserId==0 || req.FeedbackId==0 || req.Description==""{
		return nil,status.Errorf(codes.FailedPrecondition,"PLease enter all the fields")
	}

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

	if req.UserId==0 || req.FeedbackId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter both userId and feedbackId")
	}

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