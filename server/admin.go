package server

import (
	"context"
	"fmt"

	"example.com/pet-project/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m *MoviesuggestionsServiceserver) AdminLogin(ctx context.Context, req *proto.AdminLoginRequest) (*proto.AdminLoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "Please enter both email and password")
	}
	id, err := m.Db.AdminLogin(req)
	if err != nil {
		return nil, err
	}
	resp := &proto.AdminLoginResponse{
		AdminId: id,
	}
	return resp, nil
}

func (m *MoviesuggestionsServiceserver) AddMovieToDatabase(ctx context.Context, req *proto.AddMovieToDatabaseRequest) (*proto.AddMovieToDatabaseResponse, error) {

	if req.Name == "" || req.Imageurl == "" || req.Director == "" || req.Rating == 0 || req.Description == "" || req.ReleaseDate == "" || req.Movieott == "" || req.AdminId == 0 || req.Category == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "Please enter all the fields")
	}

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

	if req.MovieId == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "Please Enter Id of the movie to be deleted")
	}

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
	fmt.Println("I am first")
	if req.AdminId == 0 {
		return status.Errorf(codes.FailedPrecondition, "Please enter adminId")
	}
	fmt.Println("I am debugging this!!")
	feedbacks, err := m.Db.GetFeedBack(req)
	if err != nil {
		return err
	}
	fmt.Println("I am debugging after db getfeedback call!!")

	for _, feedback := range feedbacks {
		if err := stream.Send(&proto.GetFeedBackResponse{Description: feedback}); err != nil {
			return status.Errorf(codes.Internal, "Failed to send feedback: %v", err)
		}
	}
	return nil
}

func (m *MoviesuggestionsServiceserver) GetFeeedBack(ctx context.Context,req *proto.GetFeeedBackRequest) (*proto.GetFeeedBackResponse,error){
	if req.AdminId==0{
		return nil,status.Errorf(codes.FailedPrecondition,"Please enter AdminId")
	}

	var feedBacks []string
	feedBacks,err := m.Db.GetFeeedBack(req)
	if err!=nil{
		return nil,err
	}

	resp := &proto.GetFeeedBackResponse{}

	for _,descp := range feedBacks{
		f := descp
		resp.Description = append(resp.Description,f)
	}

	return resp,nil
}
