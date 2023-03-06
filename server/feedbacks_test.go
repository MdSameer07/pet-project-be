package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestGiveFeedBack(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().GiveFeedBack(gomock.Any()).Return(&database.FeedBack{
		User_Id: 1,
		Description: "Website is working Great",
	},nil)
	expected := &proto.GiveFeedBackResponse{
		Feedback: &proto.FeedBack{
			UserId: 1,
			Description: "Website is working Great",
		},
	}
	got,err := mssServer.GiveFeedBack(ctx,&proto.GiveFeedBackRequest{
		UserId: 1,
		Description: "Website is working Great",
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
	}
}

func TestUpdateFeedBack(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db:mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateFeedBack(gomock.Any()).Return(&database.FeedBack{
		User_Id: 1,
		Description: "I think your backend is kinda lagging",
	},nil)
	expected := &proto.UpdateFeedBackResponse{
		Feedback: &proto.FeedBack{
			UserId: 1,
			Description: "I think your backend is kinda lagging",
		},
	}
	got,err := mssServer.UpdateFeedBack(ctx,&proto.UpdateFeedBackRequest{
		UserId: 1,
		Description: "I think your backend is kinda lagging",
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v , got : %v",expected,got)
	}
}

func TestDeleteFeedBack(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db:mockDb}
	ctx := context.Background()
	Status := uint32(200)
	mockDb.EXPECT().DeleteFeedBack(gomock.Any()).Return(Status,nil)
	expected := &proto.DeleteFeedBackResponse{
		Status: 200,
		Errors: "",
	}
	got,err := mssServer.DeleteFeedBack(ctx,&proto.DeleteFeedBackRequest{
		UserId: 1,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function perfomed unexpectedly , Expected : %v , got : %v", expected,got)
	}
}