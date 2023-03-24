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

// func TestMarkAsRead(t *testing.T){
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().MarkAsRead(gomock.Any()).Return(&database.Viewed{
// 		User_Id: 2,
// 		Movie_Id: 3,
// 	},nil)
// 	expected := &proto.MarkAsReadResponse{
// 		Viewed: &proto.Viewed{
// 			UserId: 2,
// 			MovieId: 3,
// 		},
// 	}
// 	got,err := mssServer.MarkAsRead(ctx,&proto.MarkAsReadRequest{
// 		UserId: 2,
// 		MovieId: 3,
// 	})
// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The functions performed Unexpectedly , Expected : %v , got : %v",expected,got)
// 	}
// }

func TestMarkAsRead(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.MarkAsReadRequest
		mockFunc       func()
		expectedOutput *proto.MarkAsReadResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.MarkAsReadRequest{
				UserId:  3,
				MovieId: 4,
			},
			mockFunc: func() {
				mockDb.EXPECT().MarkAsRead(gomock.Any()).Return(&database.Viewed{
					User_Id:  3,
					Movie_Id: 4,
				}, nil)
			},
			expectedOutput: &proto.MarkAsReadResponse{
				Viewed: &proto.Viewed{
					UserId:  3,
					MovieId: 4,
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.MarkAsReadRequest{
				MovieId: 4,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Please enter Both UserId and MovieId"),
		},
		{
			name : "Failing testcase-2",
			input : &proto.MarkAsReadRequest{
				UserId: 3,
				MovieId: 22,
			},
			mockFunc: func() {
				mockDb.EXPECT().MarkAsRead(gomock.Any()).Return(nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.MarkAsRead(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}

// func TestMarkAsUnread(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	Status := uint32(200)
// 	mockDb.EXPECT().MarkAsUnread(gomock.Any()).Return(Status, nil)
// 	expected := &proto.MarkAsUnreadResponse{
// 		Status: 200,
// 		Errors: "",
// 	}
// 	got, err := mssServer.MarkAsUnread(ctx, &proto.MarkAsUnreadRequest{
// 		UserId:  2,
// 		MovieId: 3,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly , Expected : %v , got : %v", expected, got)
// 	}
// }

func TestMarkAsUnread(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.MarkAsUnreadRequest
		mockFunc func()
		expectedOutput *proto.MarkAsUnreadResponse
		expectedError error
	}{
		{
			name : "Passing Testcase",
			input: &proto.MarkAsUnreadRequest{
				UserId: 2,
				MovieId: 1,
			},
			mockFunc: func ()  {
				mockDb.EXPECT().MarkAsUnread(gomock.Any()).Return(uint32(200),nil)
			},
			expectedOutput: &proto.MarkAsUnreadResponse{
				Status: 200,
				Errors: "",
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.MarkAsUnreadRequest{
				MovieId: 4,
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Please enter both UserId and MovieId"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.MarkAsUnreadRequest{
				UserId: 1,
				MovieId: 25,
			},
			mockFunc: func() {
				mockDb.EXPECT().MarkAsUnread(gomock.Any()).Return(uint32(500),status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"))
			},
			expectedOutput: &proto.MarkAsUnreadResponse{
				Status: 500,
				Errors: "rpc error: code = Canceled desc = Movie with provided Id doesn't exist in the Movies Table",
			},
			expectedError: status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.MarkAsUnread(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q", got, test.expectedError)
			}
		})
	}
}
