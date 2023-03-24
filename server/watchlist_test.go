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

// func TestAddMovieToWatchList(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().AddMovieToWatchList(gomock.Any()).Return(&database.WatchList{
// 		User_Id: 3,
// 		Movie_Id: 4,
// 	},nil)
// 	expected := &proto.AddMovieToWatchListResponse{
// 		Watchlist: &proto.WatchList{
// 			UserId: 3,
// 			MovieId: 4,
// 		},
// 	}
// 	got,err := mssServer.AddMovieToWatchList(ctx,&proto.AddMovieToWatchListRequest{
// 		UserId: 3,
// 		MovieId: 4,
// 	})
// 	if err != nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The functions performed Unexpectedly, Expected : %v , Got : %v",expected,got)
// 	}
// }

func TestAddMovieToWatchList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.AddMovieToWatchListRequest
		mockFunc       func()
		expectedOutput *proto.AddMovieToWatchListResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.AddMovieToWatchListRequest{
				UserId:  3,
				MovieId: 4,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddMovieToWatchList(gomock.Any()).Return(&database.WatchList{
					User_Id:  3,
					Movie_Id: 4,
				}, nil)
			},
			expectedOutput: &proto.AddMovieToWatchListResponse{
				Watchlist: &proto.WatchList{
					UserId:  3,
					MovieId: 4,
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.AddMovieToWatchListRequest{
				MovieId: 4,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition,"Please enter UserId and MovieId"),
		},
		{
			name : "Failing testcase-2",
			input : &proto.AddMovieToWatchListRequest{
				UserId: 3,
				MovieId: 22,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddMovieToWatchList(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.AddMovieToWatchList(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}

// func TestRemoveMovieFromWatchList(t *testing.T){
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
// 	ctx := context.Background()
// 	Status := uint32(200)
// 	mockDb.EXPECT().RemoveMovieFromWatchList(gomock.Any()).Return(Status,nil)
// 	expected := &proto.RemoveMovieFromWatchListResponse{
// 		Status: 200,
// 		Errors: "",
// 	}
// 	got,err := mssServer.RemoveMovieFromWatchList(ctx,&proto.RemoveMovieFromWatchListRequest{
// 		UserId: 4,
// 		MovieId: 5,
// 	})
// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The function performed unexpectedly , Expected : %v , Got : %v",expected,got)
// 	}
// }

func TestRemoveMovieFromWatchList(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.RemoveMovieFromWatchListRequest
		mockFunc func()
		expectedOutput *proto.RemoveMovieFromWatchListResponse
		expectedError error
	}{
		{
			name : "Passing Testcase",
			input: &proto.RemoveMovieFromWatchListRequest{
				UserId: 2,
				MovieId: 1,
			},
			mockFunc: func ()  {
				mockDb.EXPECT().RemoveMovieFromWatchList(gomock.Any()).Return(uint32(200),nil)
			},
			expectedOutput: &proto.RemoveMovieFromWatchListResponse{
				Status: 200,
				Errors: "",
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.RemoveMovieFromWatchListRequest{
				MovieId: 4,
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Please enter both UserId and MovieId"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.RemoveMovieFromWatchListRequest{
				UserId: 1,
				MovieId: 25,
			},
			mockFunc: func() {
				mockDb.EXPECT().RemoveMovieFromWatchList(gomock.Any()).Return(uint32(500),status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: &proto.RemoveMovieFromWatchListResponse{
				Status: 500,
				Errors: "rpc error: code = Canceled desc = Movie with provided Id doesn't exist in the Movies Table",
			},
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.RemoveMovieFromWatchList(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}
