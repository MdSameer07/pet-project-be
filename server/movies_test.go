package server

import (
	"context"
	"io"
	"log"
	"reflect"
	"testing"
	"time"

	"example.com/pet-project/database"
	"example.com/pet-project/gen/proto"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func TestGetAllMovies(t *testing.T) {

// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	ctx := context.Background()
// 	client, close := MockServer(mockDb, ctx)
// 	defer close()

// 	mockDb.EXPECT().GetAllMovies(gomock.Any()).Return([]*proto.Movie{
// 		{
// 			Name:        "The Lord Of The Rings",
// 			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
// 			Director:    "Peter Jackson",
// 			Rating:      9.0,
// 			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
// 			ReleaseDate: "2-6-2004",
// 			Ott:         "https://www.netflix.com/in/",
// 			CategoryId:  3,
// 			AdminId:     2,
// 		},
// 		{
// 			Name:        "The Lord Of The Rings",
// 			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
// 			Director:    "Peter Jackson",
// 			Rating:      9.0,
// 			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
// 			ReleaseDate: "2-6-2004",
// 			Ott:         "https://www.netflix.com/in/",
// 			CategoryId:  3,
// 			AdminId:     2,
// 		},
// 	}, nil)

// 	expected := []*proto.Movie{
// 		{
// 			Name:        "The Lord Of The Rings",
// 			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
// 			Director:    "Peter Jackson",
// 			Rating:      9.0,
// 			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
// 			ReleaseDate: "2-6-2004",
// 			Ott:         "https://www.netflix.com/in/",
// 			CategoryId:  3,
// 			AdminId:     2,
// 		},
// 		{
// 			Name:        "The Lord Of The Rings",
// 			Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
// 			Director:    "Peter Jackson",
// 			Rating:      9.0,
// 			Description: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
// 			ReleaseDate: "2-6-2004",
// 			Ott:         "https://www.netflix.com/in/",
// 			CategoryId:  3,
// 			AdminId:     2,
// 		},
// 	}

// 	stream, err := client.GetAllMovies(ctx, &proto.GetAllMoviesRequest{})
// 	if err != nil {
// 		log.Printf("Error while getting all movies from the virtual client")
// 	}
// 	for _, expected := range expected {
// 		got, err := stream.Recv()
// 		if err == io.EOF {
// 			log.Printf("Finished")
// 			return
// 		}
// 		if err != nil {
// 			log.Printf("Error while getting a particular movie from the virtual client")
// 		}
// 		if !reflect.DeepEqual(got.Movie, expected) {
// 			t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
// 		}
// 	}
// }

func TestGetAllMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx := context.Background()
	client, close := MockServer(mockDb, ctx)
	defer close()
	tests := []struct {
		name           string
		input          *proto.GetAllMoviesRequest
		mockFunc       func()
		expectedOutput []*proto.Movie
		expectedError  error
	}{
		{
			name:  "Passing Testcase",
			input: &proto.GetAllMoviesRequest{},
			mockFunc: func() {
				mockDb.EXPECT().GetAllMovies(gomock.Any()).Return([]*proto.Movie{
					&proto.Movie{
						Name:        "The Dark Knight",
						Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
						Director:    "Cristopher Nolan",
						Rating:      9,
						CategoryId:  1,
						Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
						ReleaseDate: "18-07-2008",
						Ott:         "https://www.primevideo.com/",
						AdminId:     1,
					},
					&proto.Movie{
						Name:        "The Exorcist",
						Image:       "https://prd-rteditorial.s3.us-west-2.amazonaws.com/wp-content/uploads/2020/10/28144852/Scariest_Movies_Exorcist.jpg",
						Director:    "John Boorman",
						Rating:      8.1,
						CategoryId:  2,
						Description: "When a teenage girl is possessed by a mysterious entity, her mother seeks the help of two priests to save her daughter. A visiting actress in Washington, D.C., notices dramatic and dangerous changes in the behavior and physical make-up of her 12-year-old daughter.",
						ReleaseDate: "26-12-1973",
						Ott:         "https://www.hulu.com/",
						AdminId:     1,
					},
				}, nil)
			},
			expectedOutput: []*proto.Movie{
				&proto.Movie{
					Name:        "The Dark Knight",
					Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
					Director:    "Cristopher Nolan",
					Rating:      9,
					CategoryId:  1,
					Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
					ReleaseDate: "18-07-2008",
					Ott:         "https://www.primevideo.com/",
					AdminId:     1,
				},
				&proto.Movie{
					Name:        "The Exorcist",
					Image:       "https://prd-rteditorial.s3.us-west-2.amazonaws.com/wp-content/uploads/2020/10/28144852/Scariest_Movies_Exorcist.jpg",
					Director:    "John Boorman",
					Rating:      8.1,
					CategoryId:  2,
					Description: "When a teenage girl is possessed by a mysterious entity, her mother seeks the help of two priests to save her daughter. A visiting actress in Washington, D.C., notices dramatic and dangerous changes in the behavior and physical make-up of her 12-year-old daughter.",
					ReleaseDate: "26-12-1973",
					Ott:         "https://www.hulu.com/",
					AdminId:     1,
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			stream, err := client.GetAllMovies(ctx, test.input)
			if err != nil {
				if !reflect.DeepEqual(err, test.expectedError) {
					t.Errorf("Function perfomed Unexpectedly Got : %q, Expected : %q", err, test.expectedError)
				}
			}
			for _, expected := range test.expectedOutput {
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
		})
	}
}

func TestGetAllMoviess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	tests := []struct {
		name           string
		input          *proto.GetAllMoviessRequest
		mockFunc       func()
		expectedOutput *proto.GetAllMoviessResponse
		expectedError  error
	}{
		{
			name:  "Passing testcase",
			input: &proto.GetAllMoviessRequest{},
			mockFunc: func() {
				mockDb.EXPECT().GetAllMoviess(gomock.Any()).Return([]*proto.Movie{
					&proto.Movie{
						Name:        "The Dark Knight",
						Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
						Director:    "Cristopher Nolan",
						Rating:      9,
						CategoryId:  1,
						Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
						ReleaseDate: "18-07-2008",
						Ott:         "https://www.primevideo.com/",
						AdminId:     1,
					},
					&proto.Movie{
						Name:        "The Exorcist",
						Image:       "https://prd-rteditorial.s3.us-west-2.amazonaws.com/wp-content/uploads/2020/10/28144852/Scariest_Movies_Exorcist.jpg",
						Director:    "John Boorman",
						Rating:      8.1,
						CategoryId:  2,
						Description: "When a teenage girl is possessed by a mysterious entity, her mother seeks the help of two priests to save her daughter. A visiting actress in Washington, D.C., notices dramatic and dangerous changes in the behavior and physical make-up of her 12-year-old daughter.",
						ReleaseDate: "26-12-1973",
						Ott:         "https://www.hulu.com/",
						AdminId:     1,
					},
				}, nil)
			},
			expectedOutput: &proto.GetAllMoviessResponse{
				Movie: []*proto.Movie{
					&proto.Movie{
						Name:        "The Dark Knight",
						Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
						Director:    "Cristopher Nolan",
						Rating:      9,
						CategoryId:  1,
						Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
						ReleaseDate: "18-07-2008",
						Ott:         "https://www.primevideo.com/",
						AdminId:     1,
					},
					&proto.Movie{
						Name:        "The Exorcist",
						Image:       "https://prd-rteditorial.s3.us-west-2.amazonaws.com/wp-content/uploads/2020/10/28144852/Scariest_Movies_Exorcist.jpg",
						Director:    "John Boorman",
						Rating:      8.1,
						CategoryId:  2,
						Description: "When a teenage girl is possessed by a mysterious entity, her mother seeks the help of two priests to save her daughter. A visiting actress in Washington, D.C., notices dramatic and dangerous changes in the behavior and physical make-up of her 12-year-old daughter.",
						ReleaseDate: "26-12-1973",
						Ott:         "https://www.hulu.com/",
						AdminId:     1,
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.GetAllMoviess(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}

// func TestSearchForMovies(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	ctx := context.Background()
// 	client, close := MockServer(mockDb, ctx)
// 	defer close()
// 	mockDb.EXPECT().SearchForMovies(gomock.Any()).Return([]*proto.Movie{
// 		{
// 			Name:        "Home Alone",
// 			Image:       "https://www.dkoding.in/wp-content/uploads/Home-Alone-Part-6-Release-Date-Hollywood-Entertainment-DKODING.jpg",
// 			Director:    "Chris Columbus",
// 			Rating:      7.6,
// 			Description: "When bratty 8-year-old Kevin McCallister (Macaulay Culkin) acts out the night before a family trip to Paris, his mother (Catherine O'Hara) makes him sleep in the attic. After the McCallisters mistakenly leave for the airport without Kevin, he awakens to an empty house and assumes his wish to have no family has come true. But his excitement sours when he realizes that two con men (Joe Pesci, Daniel Stern) plan to rob the McCallister residence, and that he alone must protect the family home.",
// 			ReleaseDate: "18-10-1991",
// 			Ott:         "https://www.hotstar.com/in",
// 			CategoryId:  4,
// 			AdminId:     2,
// 		},
// 		{
// 			Name:        "The Menu",
// 			Image:       "https://i.ytimg.com/vi/LpUbk70YZd8/movieposter_en.jpg",
// 			Director:    "Mark Mylod",
// 			Rating:      7.2,
// 			Description: "The film, penned by Will Tracy and Seth Reiss, 'focuses on a young couple who visits an exclusive destination restaurant on a remote island where the acclaimed chef has prepared a lavish tasting menu, along with some shocking surprise.'Deadline notes, 'Fiennes plays the world-class chef who sets it all up and adds some unexpected ingredients to the menu planned. The action follows one particular A-list couple that takes part. I've heard Stone will play half of that couple.'",
// 			ReleaseDate: "18-11-2022",
// 			Ott:         "https://www.hotstar.com/in",
// 			CategoryId:  4,
// 			AdminId:     1,
// 		},
// 	}, nil)

// 	expected := []*proto.Movie{
// 		{
// 			Name:        "Home Alone",
// 			Image:       "https://www.dkoding.in/wp-content/uploads/Home-Alone-Part-6-Release-Date-Hollywood-Entertainment-DKODING.jpg",
// 			Director:    "Chris Columbus",
// 			Rating:      7.6,
// 			Description: "When bratty 8-year-old Kevin McCallister (Macaulay Culkin) acts out the night before a family trip to Paris, his mother (Catherine O'Hara) makes him sleep in the attic. After the McCallisters mistakenly leave for the airport without Kevin, he awakens to an empty house and assumes his wish to have no family has come true. But his excitement sours when he realizes that two con men (Joe Pesci, Daniel Stern) plan to rob the McCallister residence, and that he alone must protect the family home.",
// 			ReleaseDate: "18-10-1991",
// 			Ott:         "https://www.hotstar.com/in",
// 			CategoryId:  4,
// 			AdminId:     2,
// 		},
// 		{
// 			Name:        "The Menu",
// 			Image:       "https://i.ytimg.com/vi/LpUbk70YZd8/movieposter_en.jpg",
// 			Director:    "Mark Mylod",
// 			Rating:      7.2,
// 			Description: "The film, penned by Will Tracy and Seth Reiss, 'focuses on a young couple who visits an exclusive destination restaurant on a remote island where the acclaimed chef has prepared a lavish tasting menu, along with some shocking surprise.'Deadline notes, 'Fiennes plays the world-class chef who sets it all up and adds some unexpected ingredients to the menu planned. The action follows one particular A-list couple that takes part. I've heard Stone will play half of that couple.'",
// 			ReleaseDate: "18-11-2022",
// 			Ott:         "https://www.hotstar.com/in",
// 			CategoryId:  4,
// 			AdminId:     1,
// 		},
// 	}

// 	stream, err := client.SearchForMovies(ctx, &proto.SearchRequest{Filter: proto.SearchRequest_Category, Category: &proto.Category{Type: proto.Category_Comedy}})
// 	if err != nil {
// 		log.Printf("Error while getting all movies from the virtual cleint")
// 	}
// 	for _, expected := range expected {
// 		got, err := stream.Recv()
// 		if err == io.EOF {
// 			log.Printf("Finished")
// 			return
// 		}
// 		if err != nil {
// 			log.Printf("Error while getting a particular movie from the virtual client")
// 		}
// 		if !reflect.DeepEqual(got.Movie, expected) {
// 			t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
// 		}
// 	}
// }

func TestGetMovieById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	date := "18-07-2008"
	release_date, _ := time.Parse("02-01-2006", date)
	tests := []struct {
		name           string
		input          *proto.GetMovieByIdRequest
		mockFunc       func()
		expectedOutput *proto.GetMovieByIdResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.GetMovieByIdRequest{
				MovieId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetMovieById(gomock.Any()).Return(&database.Movie{
					MovieName:        "The Dark Knight",
					MovieImage:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
					MovieDirector:    "Cristopher Nolan",
					MovieRating:      9,
					MovieDescription: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
					MovieReleaseDate: release_date,
					MovieOtt:         "https://www.primevideo.com/",
					CategoryId:       1,
					AdminId:          1,
				}, nil)
			},
			expectedOutput: &proto.GetMovieByIdResponse{
				Movie: &proto.Movie{
					Name:        "The Dark Knight",
					Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
					Director:    "Cristopher Nolan",
					Rating:      9,
					Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
					ReleaseDate: "18-07-2008",
					Ott:         "https://www.primevideo.com/",
					CategoryId:  1,
					AdminId:     1,
					Category: &proto.Category{
						Type: 1,
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.GetMovieById(ctx, test.input)
			if !reflect.DeepEqual(got.Movie, test.expectedOutput.Movie) {
				t.Errorf("Function performed Unexpectedly, got : \n %q, expected : \n %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}

}

func TestSearchForMovies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx := context.Background()
	client, close := MockServer(mockDb, ctx)
	defer close()
	tests := []struct {
		name           string
		input          *proto.SearchRequest
		mockFunc       func()
		expectedOutput []*proto.Movie
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.SearchRequest{
				Filter: proto.SearchRequest_Category,
				Category: &proto.Category{
					Type: proto.Category_thriller,
				},
			},
			mockFunc: func() {
				mockDb.EXPECT().SearchForMovies(gomock.Any()).Return([]*proto.Movie{
					&proto.Movie{
						Name:        "The Dark Knight",
						Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
						Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
						CategoryId:  1,
						Rating:      9,
						Director:    "Cristopher Nolan",
						ReleaseDate: "18-7-2008",
						Ott:         "https://www.primevideo.com/",
						AdminId:     1,
					},
					&proto.Movie{
						Name:        "Inception",
						Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/09/Inception.jpg?q=50&fit=crop&w=1500&dpr=1.5",
						Description: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.",
						CategoryId:  1,
						Director:    "Cristopher Nolan",
						Rating:      8.8,
						ReleaseDate: "16-10-2010",
						Ott:         "https://www.primevideo.com/",
						AdminId:     2,
					},
				}, nil)
			},
			expectedOutput: []*proto.Movie{
				&proto.Movie{
					Name:        "The Dark Knight",
					Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
					Description: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
					CategoryId:  1,
					Rating:      9,
					Director:    "Cristopher Nolan",
					ReleaseDate: "18-7-2008",
					Ott:         "https://www.primevideo.com/",
					AdminId:     1,
				},
				&proto.Movie{
					Name:        "Inception",
					Image:       "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/09/Inception.jpg?q=50&fit=crop&w=1500&dpr=1.5",
					Description: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.",
					CategoryId:  1,
					Director:    "Cristopher Nolan",
					Rating:      8.8,
					ReleaseDate: "16-10-2010",
					Ott:         "https://www.primevideo.com/",
					AdminId:     2,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Failing Testcase-1",
			input:          &proto.SearchRequest{},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Please select a filter for the search"),
		},
		// {
		// 	name:  "Failing Testcase-2",
		// 	input: &proto.SearchRequest{Filter: proto.SearchRequest_Category, Category: &proto.Category{Type: proto.Category_Comedy}},
		// 	mockFunc: func() {
		// 		mockDb.EXPECT().SearchForMovies(gomock.Any()).Return(nil, nil)
		// 	},
		// 	expectedOutput: nil,
		// 	expectedError:  status.Errorf(codes.Canceled, "No movies found"),
		// },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			stream, err := client.SearchForMovies(ctx, test.input)
			if err != nil {
				if !reflect.DeepEqual(err, test.expectedError) {
					t.Errorf("Function perfomed Unexpectedly Got : %q, Expected : %q", err, test.expectedError)
				}
			}
			for _, expected := range test.expectedOutput {
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
		})
	}
}
