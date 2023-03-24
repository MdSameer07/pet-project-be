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

// func TestAddReviewForMovie(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().AddReviewForMovie(gomock.Any()).Return(&database.Review{
// 		User_Id: 1,
// 		Movie_Id: 2,
// 		Description: "This was really a heart whelming movie",
// 		Stars: 4,
// 	},nil)
// 	expected := &proto.AddReviewResponse{
// 		Review: &proto.Review{
// 			UserId: 1,
// 			MovieId: 2,
// 			Description: "This was really a heart whelming movie",
// 			Stars: 4,
// 		},
// 	}
// 	got,err := mssServer.AddReviewForMovie(ctx,&proto.AddReviewRequest{
// 		UserId: 1,
// 		MovieId: 2,
// 		Description: "This was really a heart whelming movie",
// 		Stars: 4,
// 	})
// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The function performed unexpectedly, Expected : %v , Got : %v",expected,got)
// 	}
// }

func TestAddReviewForMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.AddReviewRequest
		mockFunc       func()
		expectedOutput *proto.AddReviewResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.AddReviewRequest{
				UserId:      2,
				MovieId:     1,
				Stars:       4,
				Description: "It was a good movie",
			},
			mockFunc: func() {
				mockDb.EXPECT().AddReviewForMovie(gomock.Any()).Return(&database.Review{
					User_Id:     2,
					Movie_Id:    1,
					Stars:       4,
					Description: "It was a good movie",
				}, nil)
			},
			expectedOutput: &proto.AddReviewResponse{
				Review: &proto.Review{
					UserId:      2,
					MovieId:     1,
					Stars:       4,
					Description: "It was a good movie",
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.AddReviewRequest{
				UserId:      2,
				Stars:       4,
				Description: "It is a good movie",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Please enter UserId, MovieId and Description"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.AddReviewRequest{
				UserId: 2,
				MovieId: 1,
				Stars: 7,
				Description: "It is a nice movie",
			},
			mockFunc: func(){},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition, "Stars should be given in the range of 1 to 5"),
		},
		{
			name : "Failing Testcase-3",
			input: &proto.AddReviewRequest{
				UserId: 1,
				MovieId: 11,
				Stars: 3,
				Description: "Pretty average movie",
			},
			mockFunc: func() {
				mockDb.EXPECT().AddReviewForMovie(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.AddReviewForMovie(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}

// func TestUpdateReviewForMovie(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().UpdateReviewForMovie(gomock.Any()).Return(&database.Review{
// 		User_Id:     1,
// 		Movie_Id:    2,
// 		Description: "I am liking the movie more",
// 		Stars:       5,
// 	}, nil)
// 	expected := &proto.UpdateReviewResponse{
// 		Review: &proto.Review{
// 			UserId:      1,
// 			MovieId:     2,
// 			Description: "I am liking the movie more",
// 			Stars:       5,
// 		},
// 	}
// 	got, err := mssServer.UpdateReviewForMovie(ctx, &proto.UpdateReviewRequest{
// 		UserId:      1,
// 		MovieId:     2,
// 		Description: "I am liking the movie more",
// 		Stars:       5,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly, Expected : %v , got : %v", expected, got)
// 	}
// }

func TestUpdateReviewForMovie(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDB := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDB}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.UpdateReviewRequest
		mockFunc func()
		expectedOutput *proto.UpdateReviewResponse
		expectedError error
	}{
		{
			name : "Passing Testcase",
			input : &proto.UpdateReviewRequest{
				UserId: 1,
				MovieId: 2,
				Stars: 3,
				Description: "Feel bored after watching it more than 1 time",
			},
			mockFunc: func() {
				mockDB.EXPECT().UpdateReviewForMovie(gomock.Any()).Return(&database.Review{
					User_Id: 1,
					Movie_Id: 2,
					Stars: 3,
					Description: "Feel bored after watching it more than 1 time",
				},nil)
			},
			expectedOutput: &proto.UpdateReviewResponse{
				Review: &proto.Review{
					UserId: 1,
					MovieId: 2,
					Stars: 3,
					Description: "Feel bored after watching it more than 1 time",
				},
			},
			expectedError: nil,
		},
		{
			name : "Failing Testcase-1",
			input : &proto.UpdateReviewRequest{
				MovieId: 1,
				Stars: 2,
				Description: "Not at all recommended",
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Enter userId movieId and Description to update"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.UpdateReviewRequest{
				UserId: 11,
				MovieId: 2,
				Stars: 3,
				Description: "AN average movie to make time pass",
			},
			mockFunc: func() {
				mockDB.EXPECT().UpdateReviewForMovie(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.UpdateReviewForMovie(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}

// func TestDeleteReviewForMovie(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	Status := uint32(200)
// 	mockDb.EXPECT().DeleteReviewForMovie(gomock.Any()).Return(Status, nil)
// 	expected := &proto.DeleteReviewResponse{
// 		Status: 200,
// 		Errors: "",
// 	}
// 	got, err := mssServer.DeleteReviewForMovie(ctx, &proto.DeleteReviewRequest{
// 		UserId:  1,
// 		MovieId: 2,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly, Expected : %v , Got : %v", expected, got)
// 	}
// }

func TestDeleteReviewForMovie(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.DeleteReviewRequest
		mockFunc func()
		expectedOutput *proto.DeleteReviewResponse
		expectedError error
	}{
		{
			name : "Passing Testcase",
			input: &proto.DeleteReviewRequest{
				UserId: 2,
				MovieId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteReviewForMovie(gomock.Any()).Return(uint32(200),nil)
			},
			expectedOutput: &proto.DeleteReviewResponse{
				Status: 200,
				Errors: "",
			},
			expectedError: nil,
		},
		{
			name : "Failing Testcase-1",
			input : &proto.DeleteReviewRequest{
				UserId: 1,
			},
			mockFunc: func(){},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Please enter valid UserId and MovieId"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.DeleteReviewRequest{
				UserId: 14,
				MovieId: 2,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteReviewForMovie(gomock.Any()).Return(uint32(500),status.Errorf(codes.FailedPrecondition,"Please enter valid UserId and MovieId"))
			},
			expectedOutput: &proto.DeleteReviewResponse{
				Status: 500,
				Errors: "rpc error: code = FailedPrecondition desc = Please enter valid UserId and MovieId",
			},
			expectedError: status.Errorf(codes.FailedPrecondition,"Please enter valid UserId and MovieId"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.DeleteReviewForMovie(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}
