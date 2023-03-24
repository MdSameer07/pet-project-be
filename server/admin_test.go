package server

import (
	"context"
	"fmt"
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

// func TestAddMovieToDatabase(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
// 	ctx := context.Background()
// 	releaseDate := "19-10-1997"
// 	date, _ := time.Parse("02-01-2006", releaseDate)
// mockDb.EXPECT().AddMovieToDatabase(gomock.Any()).Return(&database.Movie{
// 	MovieName:        "Titanic",
// 	MovieImage:       "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
// 	MovieDirector:    "James Cameron",
// 	MovieReleaseDate: date,
// 	MovieDescription: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
// 	MovieRating:      7.9, MovieOtt: "https://www.hotstar.com/in",
// 	AdminId:    2,
// 	CategoryId: 5,
// }, nil)
// 	expected := &proto.AddMovieToDatabaseResponse{
// Movie: &proto.Movie{
// 	Name:        "Titanic",
// 	CategoryId:  5,
// 	Rating:      7.9,
// 	Director:    "James Cameron",
// 	ReleaseDate: "19-10-1997",
// 	Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
// 	Ott:         "https://www.hotstar.com/in",
// 	Image:       "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
// 	AdminId:     2,
// },
// 	}

// 	got, err := mssServer.AddMovieToDatabase(ctx, &proto.AddMovieToDatabaseRequest{
// Name:        "Titanic",
// Category:    "Romance",
// Rating:      7.9,
// ReleaseDate: "19-10-1997",
// Director:    "James Cameron",
// Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
// Movieott:    "https://www.hotstar.com/in",
// Imageurl:    "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
// AdminId:     2,
// 	})

// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}

// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
// 	}
// }

func TestAddMovieToDatabase(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()
	date1 := "19-10-1997"
	date1_in_time, _ := time.Parse("02-01-2006", date1)
	tests := []struct {
		name           string
		input          *proto.AddMovieToDatabaseRequest
		mockFunc       func()
		expectedOutput *proto.AddMovieToDatabaseResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.AddMovieToDatabaseRequest{
				Name:        "Titanic",
				Category:    "Romance",
				Rating:      7.9,
				ReleaseDate: "19-10-1997",
				Director:    "James Cameron",
				Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
				Movieott:    "https://www.hotstar.com/in",
				Imageurl:    "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
				AdminId:     2,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddMovieToDatabase(gomock.Any()).Return(&database.Movie{
					MovieName:        "Titanic",
					MovieImage:       "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg",
					MovieDirector:    "James Cameron",
					MovieReleaseDate: date1_in_time,
					MovieDescription: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.",
					MovieRating:      7.9, MovieOtt: "https://www.hotstar.com/in",
					AdminId:    2,
					CategoryId: 5,
				}, nil)
			},
			expectedOutput: &proto.AddMovieToDatabaseResponse{
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
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.AddMovieToDatabaseRequest{
				Name:        "Inception",
				Category:    "Thriller",
				Rating:      8.8,
				ReleaseDate: "16-10-2010",
				Director:    "Cristopher Nolan",
				Description: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.",
				Movieott:    "https://www.primevideo.com/",
				Imageurl:    "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/09/Inception.jpg?q=50&fit=crop&w=1500&dpr=1.5",
				AdminId:     7,
			},
			mockFunc: func() {
				mockDb.EXPECT().AddMovieToDatabase(gomock.Any()).Return(nil, status.Errorf(codes.NotFound, "Admin with this adminId doesn't exist in the admins table"))
			},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.NotFound, "Admin with this adminId doesn't exist in the admins table"),
		},
		{
			name: "Failing Testcase-2",
			input: &proto.AddMovieToDatabaseRequest{
				Name:        "Inception",
				Category:    "Thriller",
				ReleaseDate: "16-10-2010",
				Director:    "Cristopher Nolan",
				Description: "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.",
				Movieott:    "https://www.primevideo.com/",
				Imageurl:    "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/09/Inception.jpg?q=50&fit=crop&w=1500&dpr=1.5",
				AdminId:     7,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "PLease enter all the fields"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.AddMovieToDatabase(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}

// func TestDeleteMovieFromDatabase(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// mockDb := database.NewMockDatabase(controller)
// mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
// ctx := context.Background()
// 	mockDb.EXPECT().DeleteMovieFromDatabase(gomock.Any())
// 	expected := &proto.DeleteMovieFromDatabaseResponse{
// 		Status: 200,
// 		Error:  "",
// 	}
// 	got, err := mssServer.DeleteMovieFromDatabase(ctx, &proto.DeleteMovieFromDatabaseRequest{
// 		MovieId: 3,
// 	})
// 	if err != nil {
// 		t.Errorf(err.Error())
// 		return
// 	}
// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("The function performed unexpectedly , Expected : %v, Got : %v", expected, got)
// 	}

// }

func TestDeleteMovieFromDatabase(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	mssServer := &MoviesuggestionsServiceserver{Db: mockDb}
	ctx := context.Background()

	tests := []struct {
		name           string
		input          *proto.DeleteMovieFromDatabaseRequest
		mockFunc       func()
		expectedOutput *proto.DeleteMovieFromDatabaseResponse
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.DeleteMovieFromDatabaseRequest{
				MovieId: 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteMovieFromDatabase(gomock.Any()).Return(uint32(200), nil)
			},
			expectedOutput: &proto.DeleteMovieFromDatabaseResponse{
				Status: 200,
				Error:  "",
			},
			expectedError: nil,
		},
		{
			name: "Failing Testcase-1",
			input: &proto.DeleteMovieFromDatabaseRequest{
				MovieId: 11,
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteMovieFromDatabase(gomock.Any()).Return(uint32(500), status.Errorf(codes.NotFound, "Movie with given Id not found"))
			},
			expectedOutput: &proto.DeleteMovieFromDatabaseResponse{
				Status: 500,
				Error:  "rpc error: code = NotFound desc = Movie with given Id not found",
			},
			expectedError: status.Errorf(codes.NotFound, "Movie with given Id not found"),
		},
		{
			name:           "Failing Testcase-2",
			input:          &proto.DeleteMovieFromDatabaseRequest{},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Please Enter Id of the movie to be deleted"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			got, err := mssServer.DeleteMovieFromDatabase(ctx, test.input)
			if !reflect.DeepEqual(got, test.expectedOutput) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", got, test.expectedOutput)
			}
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Function performed Unexpectedly, got : %q, expected : %q", err, test.expectedError)
			}
		})
	}
}

// func TestGetFeedBack(t *testing.T) {

// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := database.NewMockDatabase(controller)
// 	ctx := context.Background()
// 	client, close := MockServer(mockDb, ctx)
// 	defer close()

// 	mockDb.EXPECT().GetFeedBack(gomock.Any()).Return([]string{
// 		"UI is really Awesome",
// 		"Need some optimization while getting data fast on clicking on insights",
// 	}, nil)

// 	expected := []string{
// 		"UI is really Awesome",
// 		"Need some optimization while getting data fast on clicking on insights",
// 	}

// stream, err := client.GetFeedBack(ctx, &proto.GetFeedBackRequest{AdminEmail: "sameer@gmail.com"})
// if err != nil {
// 	log.Printf("Error while getting feedbacks from the virtual client")
// }
// for _, expected := range expected {
// 	got, err := stream.Recv()
// 	if err == io.EOF {
// 		log.Printf("Finished")
// 		return
// 	}
// 	if err != nil {
// 		log.Printf("Error while getting a particular movie from the virtual client")
// 	}
// 	if !reflect.DeepEqual(got.Description, expected) {
// 		t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
// 	}
// }
// }

func TestGetFeedBack(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := database.NewMockDatabase(controller)
	ctx := context.Background()
	client, close := MockServer(mockDb, ctx)
	defer close()
	tests := []struct {
		name           string
		input          *proto.GetFeedBackRequest
		mockFunc       func()
		expectedOutput []string
		expectedError  error
	}{
		{
			name: "Passing Testcase",
			input: &proto.GetFeedBackRequest{
				AdminId : 1,
			},
			mockFunc: func() {
				mockDb.EXPECT().GetFeedBack(gomock.Any()).Return([]string{
					"UI is really Awesome",
					"Changes need to be made to optimize backend further",
				}, nil)
			},
			expectedOutput: []string{
				"UI is really Awesome",
				"Changes need to be made to optimize backend further",
			},
			expectedError: nil,
		},
		{
			name:           "Failing Testcase-1",
			input:          &proto.GetFeedBackRequest{},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  status.Errorf(codes.FailedPrecondition, "Please enter adminId"),
		},
		// {
		// 	name: "Failing Testcase-2",
		// 	input: &proto.GetFeedBackRequest{
		// 		AdminId: 4,
		// 	},
		// 	mockFunc: func() {
		// 		mockDb.EXPECT().GetFeedBack(gomock.Any()).Return(nil, status.Errorf(codes.Canceled, "Admin with provided Id doesn't exist in the Admin Table"))
		// 	},
		// 	expectedOutput: nil,
		// 	expectedError:  status.Errorf(codes.Canceled, "Admin with provided Id doesn't exist in the Admin Table"),
		// },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			fmt.Println("running", test.name)
			stream, err := client.GetFeedBack(ctx, test.input)
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
				if !reflect.DeepEqual(got.Description, expected) {
					t.Errorf("function performed Unexpectedly , Got : %v , expected : %v", got, expected)
				}
			}
		})
	}
}
