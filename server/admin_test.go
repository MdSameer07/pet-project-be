package server

import (
	"context"
	"reflect"
	"testing"
	"time"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestAddMovieToDatabase(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	releaseDate := "19-10-1997"
	date,_ := time.Parse("02-01-2006", releaseDate)
	mockDb.EXPECT().AddMovieToDatabase(gomock.Any()).Return(&database.Movie{
		MovieName:        "Titanic",
		MovieImage:       "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
		MovieDirector:    "James Cameron",
		MovieReleaseDate: date,
		MovieDescription: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
		MovieRating:      7.9, MovieOtt: "https://www.hotstar.com/in",
		AdminId:          2,
		CategoryId:       5,
	}, nil)
	expected := &proto.AddMovieToDatabaseResponse{
		Movie: &proto.Movie{
			Name:        "Titanic",
			CategoryId:  5,
			Rating:      7.9,
			Director:    "James Cameron",
			ReleaseDate: "19-10-1997",
			Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
			Ott:         "https://www.hotstar.com/in",
			Image:       "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
			AdminId:     2,
		},
	}

	got, err := mssServer.AddMovieToDatabase(ctx, &proto.AddMovieToDatabaseRequest{
		Name:        "Titanic",
		Category:    "Romance",
		Rating:      7.9,
		ReleaseDate: "19-10-1997",
		Director:    "James Cameron",
		Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
		Movieott:    "https://www.hotstar.com/in",
		Imageurl:    "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
		AdminId:     2,
	})

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
	}
}

func TestDeleteMovieFromDatabase(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db:mockDb}
	ctx := context.Background()
	mockDb.EXPECT().DeleteMovieFromDatabase(gomock.Any())
	expected := &proto.DeleteMovieFromDatabaseResponse{
		Status: 200,
		Error: "",
	}
	got,err := mssServer.DeleteMovieFromDatabase(ctx,&proto.DeleteMovieFromDatabaseRequest{
		MovieId: 3,
	})
	if err!=nil{
		t.Errorf(err.Error())
		return
	}
	if !reflect.DeepEqual(got,expected){
		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
	}

}
