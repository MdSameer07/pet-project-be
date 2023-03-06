package server

import (
	"example.com/pet-project/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) GetAllMovies(req *proto.GetAllMoviesRequest, stream proto.MovieSuggestionsService_GetAllMoviesServer) error {

	movies,err := m.Db.GetAllMovies(req)
	if err!=nil{
		return err
	}

	for _, movie := range movies {
		if err := stream.Send(&proto.GetAllMoviesResponse{Movie: movie}); err != nil {
			return status.Errorf(codes.Internal, "Failed to send movie: %v", err)
		}
	}
	return nil
}

func (m *MoviesuggestionsServiceserver) SearchForMovies(req *proto.SearchRequest, stream proto.MovieSuggestionsService_SearchForMoviesServer) error {

	movies,err := m.Db.SearchForMovies(req)
	if err!=nil{
		return err
	}

	if len(movies) == 0 {
		return status.Errorf(codes.Canceled, "No movies found")
	}

	for _, movie := range movies {
		if err := stream.Send(&proto.SearchResponse{
			Movie: movie,
		}); err != nil {
			return status.Errorf(codes.Internal, "Error while sending movies in stream: %v", err)
		}
	}
	return nil
}