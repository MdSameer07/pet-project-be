package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().Register(gomock.Any()).Return(&database.User{
		UserName: "Shreekar",
		UserEmail: "shreekar69@gmail.com",
		UserPhoneNumber: "9836473263",
	},nil)
	expected := &proto.RegisterResponse{
		User : &proto.User{
			Name: "Shreekar",
			Email: "shreekar69@gmail.com",
			PhoneNumber: "9836473263",
		},
	}

	got,err := mssServer.Register(ctx,&proto.RegisterRequest{
		Name: "Shreekar",
		Email: "shreekar69@gmail.com",
		Password: "Shreekar@123",
		PhoneNumber: "9836473263",
	})

	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
	}
}

func TestLogin(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	id := uint32(1)
	mockDb.EXPECT().Login(gomock.Any()).Return(id,nil)
	expected := &proto.LoginResponse{
		Id: 1,
	}
	got,err := mssServer.Login(ctx,&proto.LoginRequest{
		Email: "shreekar69@gmail.com",
		Password: "Shreekar@123",
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
	}
}