package server

import (
	"context"
	"io"
	"log"
	"reflect"
	"testing"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/golang/mock/gomock"
)

func TestGetAllMovies(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx := context.Background()
	client, close := MockServer(mockDb, ctx)
	defer close()

	mockDb.EXPECT().GetAllMovies(gomock.Any()).Return([]*proto.Movie{
		{
			Name:        "The Lord Of The Rings",
			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
			Director:    "Peter Jackson",
			Rating:      9.0,
			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
			ReleaseDate: "2-6-2004",
			Ott:         "https://www.netflix.com/in/",
			CategoryId:  3,
			AdminId:     2,
		},
		{
			Name:        "The Lord Of The Rings",
			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
			Director:    "Peter Jackson",
			Rating:      9.0,
			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
			ReleaseDate: "2-6-2004",
			Ott:         "https://www.netflix.com/in/",
			CategoryId:  3,
			AdminId:     2,
		},
	}, nil)

	expected := []*proto.Movie{
		{
			Name:        "The Lord Of The Rings",
			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
			Director:    "Peter Jackson",
			Rating:      9.0,
			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
			ReleaseDate: "2-6-2004",
			Ott:         "https://www.netflix.com/in/",
			CategoryId:  3,
			AdminId:     2,
		},
		{
			Name:        "The Lord Of The Rings",
			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
			Director:    "Peter Jackson",
			Rating:      9.0,
			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
			ReleaseDate: "2-6-2004",
			Ott:         "https://www.netflix.com/in/",
			CategoryId:  3,
			AdminId:     2,
		},
	}

	stream, err := client.GetAllMovies(ctx, &proto.GetAllMoviesRequest{})
	if err != nil {
		log.Printf("Error while getting all movies from the virtual cleint")
	}
	for _, expected := range expected {
		got, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Finished")
			return
		}
		if err != nil {
			log.Printf("Error while getting a particular movie from the virtual client")
		}
		if !reflect.DeepEqual(got.Movie, expected) {
			t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
		}
	}
}

func TestSearchForMovies(t *testing.T){
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx := context.Background()
	client,close := MockServer(mockDb,ctx)
	defer close()
	mockDb.EXPECT().SearchForMovies(gomock.Any()).Return([]*proto.Movie{
		{
			Name:        "Home Alone",
			Image:       "https://www.dkoding.in/wp-content/uploads/Home-Alone-Part-6-Release-Date-Hollywood-Entertainment-DKODING.jpg",
			Director:    "Chris Columbus",
			Rating:      7.6,
			Description: "When bratty 8-year-old Kevin McCallister (Macaulay Culkin) acts out the night before a family trip to Paris, his mother (Catherine O'Hara) makes him sleep in the attic. After the McCallisters mistakenly leave for the airport without Kevin, he awakens to an empty house and assumes his wish to have no family has come true. But his excitement sours when he realizes that two con men (Joe Pesci, Daniel Stern) plan to rob the McCallister residence, and that he alone must protect the family home.",
			ReleaseDate: "18-10-1991",
			Ott:         "https://www.hotstar.com/in",
			CategoryId:  4,
			AdminId:     2,
		},
		{
			Name:        "The Menu",
			Image:       "https://i.ytimg.com/vi/LpUbk70YZd8/movieposter_en.jpg",
			Director:    "Mark Mylod",
			Rating:      7.2,
			Description: "The film, penned by Will Tracy and Seth Reiss, 'focuses on a young couple who visits an exclusive destination restaurant on a remote island where the acclaimed chef has prepared a lavish tasting menu, along with some shocking surprise.'Deadline notes, 'Fiennes plays the world-class chef who sets it all up and adds some unexpected ingredients to the menu planned. The action follows one particular A-list couple that takes part. I've heard Stone will play half of that couple.'",
			ReleaseDate: "18-11-2022",
			Ott:         "https://www.hotstar.com/in",
			CategoryId:  4,
			AdminId:     1,
		},
	}, nil)

	expected := []*proto.Movie{
		{
			Name:        "Home Alone",
			Image:       "https://www.dkoding.in/wp-content/uploads/Home-Alone-Part-6-Release-Date-Hollywood-Entertainment-DKODING.jpg",
			Director:    "Chris Columbus",
			Rating:      7.6,
			Description: "When bratty 8-year-old Kevin McCallister (Macaulay Culkin) acts out the night before a family trip to Paris, his mother (Catherine O'Hara) makes him sleep in the attic. After the McCallisters mistakenly leave for the airport without Kevin, he awakens to an empty house and assumes his wish to have no family has come true. But his excitement sours when he realizes that two con men (Joe Pesci, Daniel Stern) plan to rob the McCallister residence, and that he alone must protect the family home.",
			ReleaseDate: "18-10-1991",
			Ott:         "https://www.hotstar.com/in",
			CategoryId:  4,
			AdminId:     2,
		},
		{
			Name:        "The Menu",
			Image:       "https://i.ytimg.com/vi/LpUbk70YZd8/movieposter_en.jpg",
			Director:    "Mark Mylod",
			Rating:      7.2,
			Description: "The film, penned by Will Tracy and Seth Reiss, 'focuses on a young couple who visits an exclusive destination restaurant on a remote island where the acclaimed chef has prepared a lavish tasting menu, along with some shocking surprise.'Deadline notes, 'Fiennes plays the world-class chef who sets it all up and adds some unexpected ingredients to the menu planned. The action follows one particular A-list couple that takes part. I've heard Stone will play half of that couple.'",
			ReleaseDate: "18-11-2022",
			Ott:         "https://www.hotstar.com/in",
			CategoryId:  4,
			AdminId:     1,
		},
	}

	stream, err := client.SearchForMovies(ctx,&proto.SearchRequest{Filter: proto.SearchRequest_Category,Category: &proto.Category{Type: proto.Category_Comedy}})
	if err != nil {
		log.Printf("Error while getting all movies from the virtual cleint")
	}
	for _, expected := range expected {
		got, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Finished")
			return
		}
		if err != nil {
			log.Printf("Error while getting a particular movie from the virtual client")
		}
		if !reflect.DeepEqual(got.Movie, expected) {
			t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
		}
	}
}
