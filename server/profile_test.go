package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestUpdateProfile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateProfile(gomock.Any()).Return(&database.User{
		UserName: "Shreekar",
		UserEmail: "shreekar69@gmail.com",
		UserPhoneNumber: "9999999999",
	},nil)
	expected := &proto.UpdateProfileResponse{
		Name: "Shreekar",
		Email : "shreekar69@gmail.com",
		PhoneNumber: "9999999999",
	}
	got,err := mssServer.UpdateProfile(ctx,&proto.UpdateProfileRequest{
		Id: 1,
		Name: "Shreekar",
		Email: "shreekar69@gmail.com",
		PhoneNumber: "9999999999",
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly, Expected : %v, got : %v",expected,got)
	}
}
