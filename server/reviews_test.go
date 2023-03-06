package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestAddReviewForMovie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddReviewForMovie(gomock.Any()).Return(&database.Review{
		User_Id: 1,
		Movie_Id: 2,
		Description: "This was really a heart whelming movie",
		Stars: 4,
	},nil)
	expected := &proto.AddReviewResponse{
		Review: &proto.Review{
			UserId: 1,
			MovieId: 2,
			Description: "This was really a heart whelming movie",
			Stars: 4,
		},
	}
	got,err := mssServer.AddReviewForMovie(ctx,&proto.AddReviewRequest{
		UserId: 1,
		MovieId: 2,
		Description: "This was really a heart whelming movie",
		Stars: 4,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly, Expected : %v , Got : %v",expected,got)
	}
}

func TestUpdateReviewForMovie(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().UpdateReviewForMovie(gomock.Any()).Return(&database.Review{
		User_Id: 1,
		Movie_Id: 2,
		Description: "I am liking the movie more",
		Stars: 5,
	},nil)
	expected := &proto.UpdateReviewResponse{
		Review: &proto.Review{
			UserId: 1,
			MovieId: 2,
			Description: "I am liking the movie more",
			Stars: 5,
		},
	}
	got,err := mssServer.UpdateReviewForMovie(ctx,&proto.UpdateReviewRequest{
		UserId: 1,
		MovieId: 2,
		Description: "I am liking the movie more",
		Stars: 5,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly, Expected : %v , got : %v",expected,got)
	}
}

func TestDeleteReviewForMovie(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	Status := uint32(200)
	mockDb.EXPECT().DeleteReviewForMovie(gomock.Any()).Return(Status,nil)
	expected := &proto.DeleteReviewResponse{
		Status: 200,
		Errors: "",
	}
	got,err := mssServer.DeleteReviewForMovie(ctx,&proto.DeleteReviewRequest{
		UserId: 1,
		MovieId: 2,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly, Expected : %v , Got : %v",expected,got)
	}
}
