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

// func TestAddMovieToLikes(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().AddMovieToLikes(gomock.Any()).Return(&database.Likes{
// 		User_Id: 1,
// 		Movie_Id: 2,
// 	},nil)
// 	expected := &proto.AddMovieToLikesResponse{
// 		Like: &proto.Likes{
// 			UserId: 1,
// 			MovieId: 2,
// 		},
// 	}
// 	got,err := mssServer.AddMovieToLikes(ctx,&proto.AddMovieToLikesRequest{
// 		UserId: 1,
// 		MovieId: 2,
// 	})
// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The function performed unexpectedly , Expected : %v , Got : %v",expected,got)
// 	}
// }

func TestAddMovieToLikes(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.AddMovieToLikesRequest
		mockFunc       func()
		expectedOutput *proto.AddMovieToLikesResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.AddMovieToLikesRequest{
				UserId:  14,
				MovieId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddMovieToLikes(gomock.Any()).Return(&database.Likes{
					User_Id:  14,
					Movie_Id: 1,
				}, nil)
			},
			expectedOutput: &proto.AddMovieToLikesResponse{
				Like: &proto.Likes{
					UserId:  14,
					MovieId: 1,
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.AddMovieToLikesRequest{
				UserId: 14,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Enter bith UserId and MovieId"),
		},
		{
			name : "Failing Testcase-2",
			input: &proto.AddMovieToLikesRequest{
				UserId: 13,
				MovieId: 2,
			},
			mockFunc: func ()  {
				mockDb.EXPECT().AddMovieToLikes(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.AddMovieToLikes(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}

// func TestRemoveMovieFromLikes(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	Status := uint32(200)
// 	mockDb.EXPECT().RemoveMovieFromLikes(gomock.Any()).Return(Status, nil)
// 	expected := &proto.RemoveMovieFromLikesResponse{
// 		Status: 200,
// 		Errors: "",
// 	}
// 	got, err := mssServer.RemoveMovieFromLikes(ctx, &proto.RemoveMovieFromLikesRequest{
// 		UserId:  1,
// 		MovieId: 2,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed Unexpectedly , Expected : %v , got : %v", expected, got)
// 	}
// }

func TestRemoveMovieFromLikes(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.RemoveMovieFromLikesRequest
		mockFunc func()
		expectedOutput *proto.RemoveMovieFromLikesResponse
		expectedError error
	}{
		{
			name : "Passing Testcase",
			input : &proto.RemoveMovieFromLikesRequest{
				UserId: 15,
				MovieId: 3,
			},
			mockFunc: func ()  {
				mockDb.EXPECT().RemoveMovieFromLikes(gomock.Any()).Return(uint32(200),nil)
			},
			expectedOutput: &proto.RemoveMovieFromLikesResponse{
				Status: 200,
				Errors: "",
			},
			expectedError: nil,
		},
		{
			name : "Failing Testcase-1",
			input : &proto.RemoveMovieFromLikesRequest{
				UserId: 15,
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Enter UserId and MovieId"),
		},
		{
			name : "Failing Testcase-2",
			input: &proto.RemoveMovieFromLikesRequest{
				UserId: 15,
				MovieId: 15,
			},
			mockFunc: func() {
				mockDb.EXPECT().RemoveMovieFromLikes(gomock.Any()).Return(uint32(500),status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: &proto.RemoveMovieFromLikesResponse{
				Status: 500,
				Errors: "rpc error: code = Canceled desc = Movie with provided Id doesn't exist in the Movies Table",
			},
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.RemoveMovieFromLikes(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}
