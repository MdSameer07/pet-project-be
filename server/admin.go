package server

import (
	"context"

	"example.com/pet-project/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) AddMovieToDatabase(ctx context.Context, req *proto.AddMovieToDatabaseRequest) (*proto.AddMovieToDatabaseResponse, error) {

	movie, err := m.Db.AddMovieToDatabase(req)
	if err != nil {
		return nil, err
	}
	formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
	resp := &proto.AddMovieToDatabaseResponse{
		Movie: &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Rating:      movie.MovieRating,
			Description: movie.MovieDescription,
			ReleaseDate: formatted_time,
			Ott:         movie.MovieOtt,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		},
	}
	return resp, nil
}

func (m *MoviesuggestionsServiceserver) DeleteMovieFromDatabase(ctx context.Context, req *proto.DeleteMovieFromDatabaseRequest) (*proto.DeleteMovieFromDatabaseResponse, error) {

	status, err := m.Db.DeleteMovieFromDatabase(req)
	if err != nil {
		resp := &proto.DeleteMovieFromDatabaseResponse{
			Status: status,
			Error:  err.Error(),
		}

		return resp, err
	}
	resp := &proto.DeleteMovieFromDatabaseResponse{
		Status: 200,
		Error:  "",
	}
	return resp, nil
}

func (m *MoviesuggestionsServiceserver) GetFeedBack(req *proto.GetFeedBackRequest, stream proto.MovieSuggestionsService_GetFeedBackServer) error {

	feedbacks, err := m.Db.GetFeedBack(req)
	if err != nil {
		return err
	}

	for _, feedback := range feedbacks {
		if err := stream.Send(&proto.GetFeedBackResponse{Description: feedback}); err != nil {
			return status.Errorf(codes.Internal, "Failed to send feedback: %v", err)
		}
	}
	return nil
}
