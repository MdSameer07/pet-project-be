package server

import (
	"context"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) GetAllMovies(req *proto.GetAllMoviesRequest, stream proto.MovieSuggestionsService_GetAllMoviesServer) error {

	var movies []*proto.Movie

	movies, err := m.Db.GetAllMovies(req)
	if err != nil {
		return status.Errorf(codes.Aborted, "Error while getting movies: %v", err)
	}

	for _, movie := range movies {
		if err := stream.Send(&proto.GetAllMoviesResponse{Movie: movie}); err != nil {
			return status.Errorf(codes.Internal, "Failed to send movie: %v", err)
		}
	}
	return nil
}

func (m *MoviesuggestionsServiceserver) GetAllMoviess(ctx context.Context, req *proto.GetAllMoviessRequest) (*proto.GetAllMoviessResponse, error) {
	var movies []*proto.Movie

	movies, err := m.Db.GetAllMoviess(req)
	if err != nil {
		return nil, err
	}

	resp := &proto.GetAllMoviessResponse{}

	for _, movie := range movies {
		m := &proto.Movie{
			Id:          movie.Id,
			Name:        movie.Name,
			Image:       movie.Image,
			Director:    movie.Director,
			Rating:      movie.Rating,
			Description: movie.Description,
			ReleaseDate: movie.ReleaseDate,
			Ott:         movie.Ott,
			CategoryId:  movie.CategoryId,
			AdminId:     movie.AdminId,
		}
		resp.Movie = append(resp.Movie, m)
	}
	return resp, nil
}

func (m *MoviesuggestionsServiceserver) GetMovieById(ctx context.Context, req *proto.GetMovieByIdRequest) (*proto.GetMovieByIdResponse, error) {
	if req.MovieId == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Please Enter MovidId")
	}

	movie, err := m.Db.GetMovieById(req)
	if err != nil {
		return nil, err
	}
	date := movie.MovieReleaseDate
	release_date := date.Format("02-01-2006")
	moviE := &proto.Movie{
		Id:          uint32(movie.ID),
		Name:        movie.MovieName,
		Image:       movie.MovieImage,
		Director:    movie.MovieDirector,
		Rating:      movie.MovieRating,
		Description: movie.MovieDescription,
		Ott:         movie.MovieOtt,
		CategoryId:  uint32(movie.CategoryId),
		AdminId:     uint32(movie.AdminId),
		ReleaseDate: release_date,
		Category: &proto.Category{
			Type: proto.Category_Type(movie.CategoryId),
		},
	}

	resp := &proto.GetMovieByIdResponse{
		Movie: moviE,
	}
	return resp, nil
}

func (m *MoviesuggestionsServiceserver) GetMovieByCategory(ctx context.Context, req *proto.GetMovieByCategoryRequest) (*proto.GetMovieByCategoryResponse, error) {
	if req.Category.Type.String() == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "Please select a category to fetch movies")
	}

	var movies []*proto.Movie

	movies, err := m.Db.GetMovieByCategory(req)
	if err != nil {
		return nil, err
	}
	resp := &proto.GetMovieByCategoryResponse{}

	for _, movie := range movies {
		m := &proto.Movie{
			Id:          movie.Id,
			Name:        movie.Name,
			Image:       movie.Image,
			Director:    movie.Director,
			Rating:      movie.Rating,
			Description: movie.Description,
			ReleaseDate: movie.ReleaseDate,
			Ott:         movie.Ott,
			CategoryId:  movie.CategoryId,
			AdminId:     movie.AdminId,
		}
		resp.Movie = append(resp.Movie, m)
	}
	return resp, nil

}

func (m *MoviesuggestionsServiceserver) SearchForMovies(req *proto.SearchRequest, stream proto.MovieSuggestionsService_SearchForMoviesServer) error {
	if string(req.Filter) == "" {
		return status.Errorf(codes.FailedPrecondition, "Please select a filter for the search")
	}

	var movies []*proto.Movie

	movies, err := m.Db.SearchForMovies(req)
	if err != nil {
		return err
	}

	if err != nil {
		return status.Errorf(codes.Aborted, "Error while getting movies: %v", err)
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

func (m *MoviesuggestionsServiceserver) SearchForMoviess(ctx context.Context,req *proto.SearchhRequest) (*proto.SearchhResponse,error){
	if string(req.Filter)==""{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter Which type of movie to search")
	}

	var movies []*proto.Movie

	movies, err := m.Db.SearchForMoviess(req)
	if err != nil {
		return nil,err
	}

	if len(movies) == 0 {
		return nil,status.Errorf(codes.Canceled, "No movies found")
	}

	resp := &proto.SearchhResponse{}

	for _, movie := range movies {
		m := &proto.Movie{
			Id:          movie.Id,
			Name:        movie.Name,
			Image:       movie.Image,
			Director:    movie.Director,
			Rating:      movie.Rating,
			Description: movie.Description,
			ReleaseDate: movie.ReleaseDate,
			Ott:         movie.Ott,
			CategoryId:  movie.CategoryId,
			AdminId:     movie.AdminId,
		}
		resp.Movie = append(resp.Movie, m)
	}
	return resp, nil

}
