package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestAddMovieToLikes(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddMovieToLikes(gomock.Any()).Return(&database.Likes{
		User_Id: 1,
		Movie_Id: 2,
	},nil)
	expected := &proto.AddMovieToLikesResponse{
		Like: &proto.Likes{
			UserId: 1,
			MovieId: 2,
		},
	}
	got,err := mssServer.AddMovieToLikes(ctx,&proto.AddMovieToLikesRequest{
		UserId: 1,
		MovieId: 2,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v , Got : %v",expected,got)
	}
}

func TestRemoveMovieFromLikes(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	Status := uint32(200)
	mockDb.EXPECT().RemoveMovieFromLikes(gomock.Any()).Return(Status,nil)
	expected := &proto.RemoveMovieFromLikesResponse{
		Status: 200,
		Errors: "",
	}
	got,err := mssServer.RemoveMovieFromLikes(ctx,&proto.RemoveMovieFromLikesRequest{
		UserId: 1,
		MovieId: 2,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed Unexpectedly , Expected : %v , got : %v",expected,got)
	}
}
