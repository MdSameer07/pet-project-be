package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/gen/proto"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func TestGiveFeedBack(t *testing.T){
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := &MoviesuggestionsServiceserver{Db : mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().GiveFeedBack(gomock.Any()).Return(&database.FeedBack{
// 		User_Id: 1,
// 		Description: "Website is working Great",
// 	},nil)
// 	expected := &proto.GiveFeedBackResponse{
// 		Feedback: &proto.FeedBack{
// 			UserId: 1,
// 			Description: "Website is working Great",
// 		},
// 	}
// 	got,err := mssServer.GiveFeedBack(ctx,&proto.GiveFeedBackRequest{
// 		UserId: 1,
// 		Description: "Website is working Great",
// 	})
// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
// 	}
// }

func TestGiveFeedBack(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.GiveFeedBackRequest
		mockFunc       func()
		expectedOutput *proto.GiveFeedBackResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.GiveFeedBackRequest{
				UserId:      14,
				Description: "UI is really fascinating",
			},
			mockFunc: func() {
				mockDb.EXPECT().GiveFeedBack(gomock.Any()).Return(&database.FeedBack{
					User_Id:     14,
					Description: "UI is really fascinating",
				}, nil)
			},
			expectedOutput: &proto.GiveFeedBackResponse{
				Feedback: &proto.FeedBack{
					UserId:      14,
					Description: "UI is really fascinating",
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.GiveFeedBackRequest{
				UserId: 14,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Enter both userId and description of the feedback"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.GiveFeedBackRequest{
				UserId: 13,
				Description: "There's lot of lag while fetching movies from the database",
			},
			mockFunc: func ()  {
				mockDb.EXPECT().GiveFeedBack(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.GiveFeedBack(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}

}

// func TestUpdateFeedBack(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().UpdateFeedBack(gomock.Any()).Return(&database.FeedBack{
// 		User_Id:     1,
// 		Description: "I think your backend is kinda lagging",
// 	}, nil)
// 	expected := &proto.UpdateFeedBackResponse{
// 		Feedback: &proto.FeedBack{
// 			UserId:      1,
// 			Description: "I think your backend is kinda lagging",
// 		},
// 	}
// 	got, err := mssServer.UpdateFeedBack(ctx, &proto.UpdateFeedBackRequest{
// 		UserId:      1,
// 		Description: "I think your backend is kinda lagging",
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly , Expected : %v , got : %v", expected, got)
// 	}
// }

func TestUpdateFeedBack(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.UpdateFeedBackRequest
		mockFunc func()
		expectedOutput *proto.UpdateFeedBackResponse
		expectedError error
	}{
		{
			name : "Passing testcase",
			input: &proto.UpdateFeedBackRequest{
				UserId: 14,
				FeedbackId: 1,
				Description: "UI is good , but change the colors of the background",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateFeedBack(gomock.Any()).Return(&database.FeedBack{
					User_Id: 14,
					Description: "UI is good , but change the colors of the background",
				},nil)
			},
			expectedOutput: &proto.UpdateFeedBackResponse{
				Feedback: &proto.FeedBack{
					UserId: 14,
					Description: "UI is good , but change the colors of the background",
				},
			},
			expectedError: nil,
		},
		{
			name : "Failing Testcase-1",
			input: &proto.UpdateFeedBackRequest{
				UserId: 14,
				Description: "UI is good , but change the colors of the background",
			},
			mockFunc: func ()  {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"PLease enter all the fields"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.UpdateFeedBackRequest{
				UserId: 14,
				FeedbackId: 2,
				Description: "UI is good , but change the colors of the background",
			},
			mockFunc: func ()  {
				mockDb.EXPECT().UpdateFeedBack(gomock.Any()).Return(nil,status.Errorf(codes.NotFound, "User has not added this particular feedback"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.NotFound, "User has not added this particular feedback"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.UpdateFeedBack(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}

}

// func TestDeleteFeedBack(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	Status := uint32(200)
// 	mockDb.EXPECT().DeleteFeedBack(gomock.Any()).Return(Status, nil)
// 	expected := &proto.DeleteFeedBackResponse{
// 		Status: 200,
// 		Errors: "",
// 	}
// 	got, err := mssServer.DeleteFeedBack(ctx, &proto.DeleteFeedBackRequest{
// 		UserId: 1,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function perfomed unexpectedly , Expected : %v , got : %v", expected, got)
// 	}
// }

func TestDeleteFeedBack(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *proto.DeleteFeedBackRequest
		mockFunc       func()
		expectedOutput *proto.DeleteFeedBackResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.DeleteFeedBackRequest{
				UserId: 15,
				FeedbackId: 2,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteFeedBack(gomock.Any()).Return(uint32(200), nil)
			},
			expectedOutput: &proto.DeleteFeedBackResponse{
				Status: 200,
				Errors:  "",
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.DeleteFeedBackRequest{
				UserId: 17,
				FeedbackId: 3,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteFeedBack(gomock.Any()).Return(uint32(500), status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"))
			},
			expectedOutput: &proto.DeleteFeedBackResponse{
				Status: 500,
				Errors:  "rpc error: code = Canceled desc = User with provided Id doesn't exist in the User Table",
			},
			expectedError: status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"),
		},
		{
			name:           "Failing Testcase-2",
			input:          &proto.DeleteFeedBackRequest{
				UserId: 15,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition,"Please enter both userId and feedbackId"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.DeleteFeedBack(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}
