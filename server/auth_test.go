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

// func TestRegister(t *testing.T){
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	mockDb.EXPECT().Register(gomock.Any()).Return(&database.User{
// 		UserName: "Shreekar",
// 		UserEmail: "shreekar69@gmail.com",
// 		UserPhoneNumber: "9836473263",
// 	},nil)
// 	expected := &proto.RegisterResponse{
// 		User : &proto.User{
// 			Name: "Shreekar",
// 			Email: "shreekar69@gmail.com",
// 			PhoneNumber: "9836473263",
// 		},
// 	}

// 	got,err := mssServer.Register(ctx,&proto.RegisterRequest{
// 		Name: "Shreekar",
// 		Email: "shreekar69@gmail.com",
// 		Password: "Shreekar@123",
// 		PhoneNumber: "9836473263",
// 	})

// 	if err!=nil{
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got,expected){
// 		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
// 	}
// }

func TestRegister(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.RegisterRequest
		mockFunc       func()
		expectedOutput *proto.RegisterResponse
		expectedError  error
	}{
		{
			name: "Passing testcase",
			input: &proto.RegisterRequest{
				Name:        "Rohith",
				Email:       "rohoth@gmail.com",
				Password:    "Rohith@123",
				PhoneNumber: "9494949494",
			},
			mockFunc: func() {
				mockDb.EXPECT().Register(gomock.Any()).Return(&database.User{
					UserName:        "Rohith",
					UserEmail:       "rohith@gmail.com",
					UserPhoneNumber: "9494949494",
				}, nil)
			},
			expectedOutput: &proto.RegisterResponse{
				User: &proto.User{
					Name:        "Rohith",
					Email:       "rohith@gmail.com",
					PhoneNumber: "9494949494",
				},
			},
			expectedError: nil,
		},
		{
			name: "Failing testcase-1",
			input: &proto.RegisterRequest{
				Name:        "Rohith",
				Email:       "rohoth@gmail.com",
				Password:    "Rohith@123",
				PhoneNumber: "9494949494",
			},
			mockFunc: func() {
				mockDb.EXPECT().Register(gomock.Any()).Return(nil, status.Errorf(codes.AlreadyExists, "User already exists with provided email"))
			},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.AlreadyExists, "User already exists with provided email"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.RegisterRequest{
				Name: "Rohith",
				Email: "rohith@gmail.com",
				PhoneNumber: "9494949494",
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition, "Enter all the fields"),
		},
	}
	for _,test := range tests{
		t.Run(test.name,func(t *testing.T){
			test.mockFunc()
			got,err := mssServer.Register(ctx,test.input)
			if !reflect.DeepEqual(got,test.expectedOutput){
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q",got,test.expectedOutput)
			}
			if !reflect.DeepEqual(err,test.expectedError){
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q",got,test.expectedError)
			}
		})
	}
}

// func TestLogin(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	id := uint32(1)
// 	mockDb.EXPECT().Login(gomock.Any()).Return(id, nil)
// 	expected := &proto.LoginResponse{
// 		Id: 1,
// 	}
// 	got, err := mssServer.Login(ctx, &proto.LoginRequest{
// 		Email:    "shreekar69@gmail.com",
// 		Password: "Shreekar@123",
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
// 	}
// }

func TestLogin(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	tests := []struct{
		name string
		input *proto.LoginRequest
		mockFunc func()
		expectedOutput *proto.LoginResponse
		expectedError error
	}{
		{
			name: "Passing Testcase",
			input : &proto.LoginRequest{
				Email: "shreekar69@gmail.com",
				Password: "Shreekar@123",
			},
			mockFunc: func() {
				mockDb.EXPECT().Login(gomock.Any()).Return(uint32(20),nil)
			},
			expectedOutput: &proto.LoginResponse{
				Id: 20,
			},
			expectedError: nil,
		},
		{
			name : "Failing Testcase-1",
			input : &proto.LoginRequest{
				Email: "shreekar@gmail.com",
				Password: "Shreekar@123",
			},
			mockFunc: func() {
				mockDb.EXPECT().Login(gomock.Any()).Return(uint32(0),status.Errorf(codes.NotFound,"Enter valid email Id"))
			},
			expectedOutput: &proto.LoginResponse{
				Id: uint32(0),
			},
			expectedError: status.Errorf(codes.NotFound,"Enter valid email Id"),
		},
		{
			name : "Failing Testcase-2",
			input : &proto.LoginRequest{
				Email : "shreekar69@gmail.com",
				Password: "shreekar@123",
			},
			mockFunc: func() {
				mockDb.EXPECT().Login(gomock.Any()).Return(uint32(0),status.Errorf(codes.FailedPrecondition,"Enter correct password"))
			},
			expectedOutput: &proto.LoginResponse{
				Id : uint32(0),
			},
			expectedError: status.Errorf(codes.FailedPrecondition,"Enter correct password"),
		},
		{
			name: "Failing Testcase-3",
			input: &proto.LoginRequest{
				Email: "shreekar69@gmail.com",
			},
			mockFunc: func() {},
			expectedOutput: nil,
			expectedError: status.Errorf(codes.FailedPrecondition,"Enter all the fields"),
		},
	}

	for _,test := range tests{
		t.Run(test.name,func(t *testing.T){
			test.mockFunc()
			got,err := mssServer.Login(ctx,test.input)
			if !reflect.DeepEqual(got,test.expectedOutput){
				t.Errorf("The function perfomed unexpectedly, got : %q, expected : %q",got,test.expectedOutput)
			}
			if !reflect.DeepEqual(err,test.expectedError){
				t.Errorf("The function performed unexpectedly, got : %q, expected : %q",got,test.expectedError)
			}
		})
	}
}
