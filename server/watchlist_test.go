package server

import (
	"context"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestAddMovieToWatchList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	mockDb.EXPECT().AddMovieToWatchList(gomock.Any()).Return(&database.WatchList{
		User_Id: 3,
		Movie_Id: 4,
	},nil)
	expected := &proto.AddMovieToWatchListResponse{
		Watchlist: &proto.WatchList{
			UserId: 3,
			MovieId: 4,
		},
	}
	got,err := mssServer.AddMovieToWatchList(ctx,&proto.AddMovieToWatchListRequest{
		UserId: 3,
		MovieId: 4,
	})
	if err != nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The functions performed Unexpectedly, Expected : %v , Got : %v",expected,got)
	}
}

func TestRemoveMovieFromWatchList(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db : mockDb}
	ctx := context.Background()
	Status := uint32(200)
	mockDb.EXPECT().RemoveMovieFromWatchList(gomock.Any()).Return(Status,nil)
	expected := &proto.RemoveMovieFromWatchListResponse{
		Status: 200,
		Errors: "",
	}
	got,err := mssServer.RemoveMovieFromWatchList(ctx,&proto.RemoveMovieFromWatchListRequest{
		UserId: 4,
		MovieId: 5,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v , Got : %v",expected,got)
	}
}
