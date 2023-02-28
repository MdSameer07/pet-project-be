package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestMarkAsRead(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().MarkAsRead(gomock.Any()).Return(&database.Viewed{
		User_Id: 2,
		Movie_Id: 3,
	},nil)
	expected := &proto.MarkAsReadResponse{
		Viewed: &proto.Viewed{
			UserId: 2,
			MovieId: 3,
		},
	}
	got,err := mssServer.MarkAsRead(ctx,&proto.MarkAsReadRequest{
		UserId: 2,
		MovieId: 3,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The functions performed Unexpectedly , Expected : %v , got : %v",expected,got)
	}
}

func TestMarkAsUnread(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	Status := uint32(200)
	mockDb.EXPECT().MarkAsUnread(gomock.Any()).Return(Status,nil)
	expected := &proto.MarkAsUnreadResponse{
		Status: 200,
		Errors: "",
	}
	got,err := mssServer.MarkAsUnread(ctx,&proto.MarkAsUnreadRequest{
		UserId: 2,
		MovieId: 3,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v , got : %v",expected,got)
	}
}