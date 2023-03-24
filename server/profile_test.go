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

// func TestUpdateProfile(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().UpdateProfile(gomock.Any()).Return(&database.User{
// 		UserName:        "Shreekar",
// 		UserEmail:       "shreekar69@gmail.com",
// 		UserPhoneNumber: "9999999999",
// 	}, nil)
// 	expected := &proto.UpdateProfileResponse{
// 		Name:        "Shreekar",
// 		Email:       "shreekar69@gmail.com",
// 		PhoneNumber: "9999999999",
// 	}
// 	got, err := mssServer.UpdateProfile(ctx, &proto.UpdateProfileRequest{
// 		Id:          1,
// 		Name:        "Shreekar",
// 		Email:       "shreekar69@gmail.com",
// 		PhoneNumber: "9999999999",
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly, Expected : %v, got : %v", expected, got)
// 	}
// }

func TestUpdateProfile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *proto.UpdateProfileRequest
		mockFunc       func()
		expectedOutput *proto.UpdateProfileResponse
		expectedError  error
	}{
		{
			name: "Passing Testacase",
			input: &proto.UpdateProfileRequest{
				Id:          1,
				Name:        "Shreekar",
				Email:       "shreekar69@gmail.com",
				PhoneNumber: "6969696969",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateProfile(gomock.Any()).Return(&database.User{
					UserName:        "Shreekar",
					UserEmail:       "shreekar69@gmail.com",
					UserPhoneNumber: "6969696969",
				}, nil)
			},
			expectedOutput: &proto.UpdateProfileResponse{
				Name:        "Shreekar",
				Email:       "shreekar69@gmail.com",
				PhoneNumber: "6969696969",
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.UpdateProfileRequest{
				Id:          4,
				Name:        "Ravi",
				Email:       "ravi@gmail.com",
				PhoneNumber: "9494949494",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateProfile(gomock.Any()).Return(nil, status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"))
			},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.UpdateProfileRequest{
				Name: "Ravi",
				Email : "ravi@gmail.com",
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Enter id of user to be updated"),
		},
	}
	for _,test := range tests{
		t.Run(test.name,func(t *testing.T){
			test.mockFunc()
			got,err := mssServer.UpdateProfile(ctx,test.input)
			if !reflect.DeepEqual(got,test.expectedOutput){
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q",got,test.expectedOutput)
			}
			if !reflect.DeepEqual(err,test.expectedError){
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q",got,test.expectedError)
			}
		})
	}
}