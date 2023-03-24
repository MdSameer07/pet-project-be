package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"example.com/pet-project/gen/proto"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Database interface {
	AdminLogin(*proto.AdminLoginRequest) (uint32,error)
	AddMovieToDatabase(*proto.AddMovieToDatabaseRequest)(*Movie,error)
	DeleteMovieFromDatabase(*proto.DeleteMovieFromDatabaseRequest) (uint32,error)
	GetFeedBack(*proto.GetFeedBackRequest) ([]string,error)
	GetFeeedBack(*proto.GetFeeedBackRequest) ([]string,error)
	Login(*proto.LoginRequest) (uint32,error)
	Register(*proto.RegisterRequest) (*User,error)
	GiveFeedBack(*proto.GiveFeedBackRequest) (*FeedBack,error)
	UpdateFeedBack(*proto.UpdateFeedBackRequest) (*FeedBack,error)
	DeleteFeedBack(*proto.DeleteFeedBackRequest) (uint32,error)
	AddMovieToLikes(*proto.AddMovieToLikesRequest) (*Likes,error)
	RemoveMovieFromLikes(*proto.RemoveMovieFromLikesRequest) (uint32,error)
	GetAllMovies(*proto.GetAllMoviesRequest) ([]*proto.Movie,error)
	GetAllMoviess(*proto.GetAllMoviessRequest) ([]*proto.Movie,error)
	GetMovieById(*proto.GetMovieByIdRequest) (*Movie,error)
	GetMovieByCategory(*proto.GetMovieByCategoryRequest) ([]*proto.Movie,error)
	SearchForMovies(*proto.SearchRequest) ([]*proto.Movie,error)
	SearchForMoviess(*proto.SearchhRequest) ([]*proto.Movie,error)
	UpdateProfile(*proto.UpdateProfileRequest) (*User,error)
	AddReviewForMovie(*proto.AddReviewRequest) (*Review,error)
	UpdateReviewForMovie(*proto.UpdateReviewRequest) (*Review,error)
	DeleteReviewForMovie(*proto.DeleteReviewRequest) (uint32,error)
	MarkAsRead(*proto.MarkAsReadRequest) (*Viewed,error)
	MarkAsUnread(*proto.MarkAsUnreadRequest) (uint32,error)
	AddMovieToWatchList(*proto.AddMovieToWatchListRequest) (*WatchList,error)
	RemoveMovieFromWatchList(*proto.RemoveMovieFromWatchListRequest) (uint32,error)
}

type DBClient struct {
	Db *gorm.DB
}

func(db DBClient) AdminLogin(req *proto.AdminLoginRequest) (uint32,error){

	var admin Admin
	if err := db.Db.Model(&Admin{}).Where("admin_email=? and admin_password=?",req.Email,req.Password).Find(&admin).Error; err!=nil{
		return 0,status.Errorf(codes.NotFound,"Enter valid email Id or password")
	}

	return uint32(admin.ID),nil
}

func (db DBClient) AddMovieToDatabase(req *proto.AddMovieToDatabaseRequest) (*Movie,error) {
	var admin Admin
	result := db.Db.Debug().Model(&Admin{}).Where("id=?", req.GetAdminId()).Find(&admin)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil, status.Errorf(codes.NotFound, "Admin with this adminId doesn't exist in the admins table")
		}
	}

	release_date, err := time.Parse("02-01-2006", req.ReleaseDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error while converting Time into correct format: %v", err)
	}

	var category Category

	if err := db.Db.Debug().FirstOrCreate(&category, Category{Label: strings.ToLower(req.Category)}).Error; err != nil {
		return nil, status.Errorf(codes.Aborted, "Error while creating category file: %v", err)
	}

	fmt.Println(float32(req.Rating))
	movie := &Movie{
		MovieName:        req.Name,
		MovieImage:       req.Imageurl,
		MovieDirector:    req.Director,
		MovieDescription: req.Description,
		MovieRating:      float32(req.Rating),
		MovieReleaseDate: release_date,
		MovieOtt:         req.Movieott,
		AdminId:          int(admin.ID),
		CategoryId:       int(category.ID),
	}

	if err := db.Db.Debug().Create(&movie).Error; err != nil {
		return nil, status.Errorf(codes.Canceled, "Error while creating a Movie: %v", err)
	}
	fmt.Println(movie)

	return movie,nil
}

func (db DBClient) DeleteMovieFromDatabase(req *proto.DeleteMovieFromDatabaseRequest) (uint32,error){
	var movie Movie
	result := db.Db.Debug().Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "Movie with gievn Id not found")
		}
	}

	if err := db.Db.Debug().Unscoped().Delete(&movie).Error; err != nil {
		return 500, status.Errorf(codes.Internal, "Error while deleting movie from database: %v", err)
	}

	return 200,nil
}

func (db DBClient) GetFeedBack(req *proto.GetFeedBackRequest) ([]string,error){
	var admin Admin
	result := db.Db.Model(&Admin{}).Where("id=?", req.AdminId).Find(&admin)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Admin with provided Id doesn't exist in the Admin Table")
		}
	}

	var feedbacks []string
	rows, err := db.Db.Model(&FeedBack{}).Rows()
	if err != nil {
		return nil,status.Errorf(codes.Aborted, "Error while getting feedbacks from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var feedback FeedBack
		err := db.Db.ScanRows(rows, &feedback)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning feedbacks into slice: %v", err)
		}
		feedbacks = append(feedbacks, feedback.Description)
	}

	return feedbacks,nil
}

func (db DBClient) GetFeeedBack(req *proto.GetFeeedBackRequest) ([]string,error){
	var admin Admin
	result := db.Db.Model(&Admin{}).Where("id=?",req.AdminId).Find(&admin)
	if result.Error!=nil{
		if result.RecordNotFound()==true{
			return nil,status.Errorf(codes.Canceled,"Admin with provoded Id doesn't exist in the admins Table")
		}
	}

	var feedbacks []string
	rows, err := db.Db.Model(&FeedBack{}).Rows()
	if err != nil {
		return nil,status.Errorf(codes.Aborted, "Error while getting feedbacks from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var feedback FeedBack
		err := db.Db.ScanRows(rows, &feedback)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning feedbacks into slice: %v", err)
		}
		feedbacks = append(feedbacks, feedback.Description)
	}

	return feedbacks,nil
}

func (db DBClient) Login(req *proto.LoginRequest) (uint32,error){
	var user User
	if err := db.Db.Model(&User{}).Where("user_email=?",req.Email).Find(&user).Error; err!=nil{
		return 0,status.Errorf(codes.NotFound,"Enter valid email Id")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword),[]byte(req.Password))
	if err!=nil{
		return 0,status.Errorf(codes.FailedPrecondition,"Enter correct password")
	}
	return uint32(user.ID),nil
}

func (db DBClient) Register(req *proto.RegisterRequest) (*User,error){
	var count int
	db.Db.Model(&User{}).Where("user_email=?",req.Email).Count(&count)

	if count>=1{
		return nil,status.Errorf(codes.AlreadyExists,"User already exists with provided email")
	}

	HashedPassword,err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err!=nil{
		return nil,status.Errorf(codes.Canceled,"Error while hashing password")
	}

	user := &User{
		UserName: req.Name,
		UserPassword: string(HashedPassword),
		UserEmail: req.Email,
		UserPhoneNumber: req.PhoneNumber,
	}

	if err := db.Db.Create(&user).Error; err!=nil{
		return nil,status.Errorf(codes.Internal,"Error while creating a new User")
	}

	return user,nil
}

func (db DBClient) GiveFeedBack(req *proto.GiveFeedBackRequest) (*FeedBack,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	feedback := &FeedBack{
		User_Id:     int(req.UserId),
		Description: req.Description,
	}

	if err := db.Db.Create(&feedback).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error while adding a feedback")
	}
	
	return feedback,nil
}

func (db DBClient) UpdateFeedBack(req *proto.UpdateFeedBackRequest) (*FeedBack,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var feedback1 FeedBack
	result = db.Db.Model(&FeedBack{}).Where("id=?", req.FeedbackId).Find(&feedback1)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil, status.Errorf(codes.Canceled, "Feedback with provided Id doesn't exist in the feedback Table")
		}
	}

	var feedback2 FeedBack
	result = db.Db.Model(&FeedBack{}).Where("id=? and user_id=?", req.FeedbackId, req.UserId).Find(&feedback2)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil, status.Errorf(codes.NotFound, "User has not added this particular feedback")
		}
	}

	feedback2.Description = req.Description
	if err := db.Db.Save(&feedback2).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error while saving the changes in the database: %v", err)
	}

	return &feedback2,nil
}

func (db DBClient) DeleteFeedBack(req *proto.DeleteFeedBackRequest) (uint32,error){
	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var feedback1 FeedBack
	result = db.Db.Model(&FeedBack{}).Where("id=?", req.FeedbackId).Find(&feedback1)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.Canceled, "Feedback with provided Id doesn't exist in the feedback Table")
		}
	}

	var feedback2 FeedBack
	result = db.Db.Model(&FeedBack{}).Where("id=? and user_id=?", req.FeedbackId, req.UserId).Find(&feedback2)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "No particular row with given user_id and feedback_id in feedback table")
		}
	}

	if err := db.Db.Unscoped().Delete(&feedback2).Error; err != nil {
		return 500, status.Errorf(codes.Aborted, "Error while deleting a feedback from FeedBack Table")
	}

	return 200,nil
}

func (db DBClient) AddMovieToLikes(req *proto.AddMovieToLikesRequest) (*Likes,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	likes := &Likes{
		User_Id:  int(req.UserId),
		Movie_Id: int(req.MovieId),
	}

	if err := db.Db.Create(&likes).Error; err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Error while adding movie to likes: %v", err)
	}

	return likes,nil
}

func (db DBClient) RemoveMovieFromLikes(req *proto.RemoveMovieFromLikesRequest) (uint32,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	var likes Likes
	result = db.Db.Model(&Likes{}).Where("user_id=? and movie_id=?", req.UserId, req.MovieId).Find(&likes)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "User has not added this particular movie into likes")
		}
	}

	if err := db.Db.Unscoped().Delete(&likes).Error; err != nil {
		return 500, status.Errorf(codes.Aborted, "Error while deleting movie from Likes: %v", err)
	}

	return 200,nil
}

func (db DBClient) GetAllMovies(req *proto.GetAllMoviesRequest) ([]*proto.Movie,error){
	rows, err := db.Db.Model(&Movie{}).Rows()
	if err != nil {
		return nil,status.Errorf(codes.Aborted, "Error while getting movies from database: %v", err)
	}
	var movies []*proto.Movie
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := db.Db.ScanRows(rows, &movie)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
		}
		formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
		movies = append(movies, &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Description: movie.MovieDescription,
			Rating:      movie.MovieRating,
			Ott:         movie.MovieOtt,
			ReleaseDate: formatted_time,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		})
	}
	return movies,nil
}

func (db DBClient) GetAllMoviess(req *proto.GetAllMoviessRequest) ([]*proto.Movie,error){
	rows, err := db.Db.Model(&Movie{}).Rows()
	if err != nil {
		return nil,status.Errorf(codes.Aborted, "Error while getting movies from database: %v", err)
	}
	var movies []*proto.Movie
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := db.Db.ScanRows(rows, &movie)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
		}
		formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
		movies = append(movies, &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Description: movie.MovieDescription,
			Rating:      movie.MovieRating,
			Ott:         movie.MovieOtt,
			ReleaseDate: formatted_time,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		})
	}
	return movies,nil
}

func (db DBClient) GetMovieById(req *proto.GetMovieByIdRequest) (*Movie,error){
	var movie Movie
	result := db.Db.Model(&Movie{}).Preload("Category").Where("id=?",req.MovieId).Find(&movie)
	if result.Error!=nil{
		if result.RecordNotFound()==true{
			return nil,status.Errorf(codes.NotFound,"Movie with given Id not found in the Movies Table")
		}
	}
	return &movie,nil
}

func (db DBClient) GetMovieByCategory(req *proto.GetMovieByCategoryRequest) ([]*proto.Movie,error){
	var category Category
	split := strings.Split(req.Category.Type.String(), "_")
	category_name := strings.ToLower(split[len(split)-1])

	if err := db.Db.Model(&Category{}).Where("label=?",category_name).Find(&category).Error; err!=nil{
		return nil,status.Errorf(codes.Aborted,"Error while fetching movie based on category")
	}

	rows,err := db.Db.Model(&Movie{}).Where("category_id=?",category.ID).Rows()
	if err!=nil{
		return nil,status.Errorf(codes.Aborted,"Error while getting movies from database")
	}

	var movies []*proto.Movie
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := db.Db.ScanRows(rows, &movie)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
		}
		formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
		movies = append(movies, &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Description: movie.MovieDescription,
			Rating:      movie.MovieRating,
			Ott:         movie.MovieOtt,
			ReleaseDate: formatted_time,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		})
	}
	return movies,nil
}

func (db DBClient) SearchForMovies(req *proto.SearchRequest) ([]*proto.Movie,error){
	var rows *sql.Rows
	var err error
	switch req.Filter {
	case proto.SearchRequest_Name:
		rows, err = db.Db.Model(&Movie{}).Where("lower(movie_name) like ?", "%"+strings.ToLower(req.Name)+"%").Rows()
		if err != nil {
			return nil,status.Errorf(codes.Canceled, "Error while fetching movie with given name: %v", err)
		}
	case proto.SearchRequest_Category:
		var category Category
		split := strings.Split(req.Category.Type.String(),"_")
		category_name := strings.ToLower(split[len(split)-1])

		if err := db.Db.Model(&Category{}).Where("label=?",category_name).Find(&category).Error; err!=nil{
			return nil,status.Errorf(codes.Aborted,"Error while fetching movie based on category")
		}

		rows,err = db.Db.Model(&Movie{}).Where("category_id=?",category.ID).Rows()
		if err!=nil{
			return nil,status.Errorf(codes.Aborted,"Error while getting movies from database")
		}
	default:
		return nil,status.Errorf(codes.Aborted, "Error!! Enter correct name or category")
	}
	var movies []*proto.Movie
	for rows.Next() {
		var movie Movie
		err := db.Db.ScanRows(rows, &movie)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
		}
		formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
		movies = append(movies, &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Description: movie.MovieDescription,
			Rating:      movie.MovieRating,
			Ott:         movie.MovieOtt,
			ReleaseDate: formatted_time,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		})
	}
	return movies,nil
}

func (db DBClient) SearchForMoviess(req *proto.SearchhRequest) ([]*proto.Movie,error){
	var rows *sql.Rows
	var err error
	switch req.Filter {
	case proto.SearchhRequest_Name:
		rows, err = db.Db.Model(&Movie{}).Where("lower(movie_name) like ?", "%"+strings.ToLower(req.Name)+"%").Rows()
		if err != nil {
			return nil,status.Errorf(codes.Canceled, "Error while fetching movie with given name: %v", err)
		}
	case proto.SearchhRequest_Category:
		var category Category
		split := strings.Split(req.Category.Type.String(),"_")
		category_name := strings.ToLower(split[len(split)-1])

		if err := db.Db.Model(&Category{}).Where("label=?",category_name).Find(&category).Error; err!=nil{
			return nil,status.Errorf(codes.Aborted,"Error while fetching movie based on category")
		}

		rows,err = db.Db.Model(&Movie{}).Where("category_id=?",category.ID).Rows()
		if err!=nil{
			return nil,status.Errorf(codes.Aborted,"Error while getting movies from database")
		}
	default:
		return nil,status.Errorf(codes.Aborted, "Error!! Enter correct name or category")
	}
	var movies []*proto.Movie
	for rows.Next() {
		var movie Movie
		err := db.Db.ScanRows(rows, &movie)
		if err != nil {
			return nil,status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
		}
		formatted_time := movie.MovieReleaseDate.Format("02-01-2006")
		movies = append(movies, &proto.Movie{
			Id:          uint32(movie.ID),
			Name:        movie.MovieName,
			Image:       movie.MovieImage,
			Director:    movie.MovieDirector,
			Description: movie.MovieDescription,
			Rating:      movie.MovieRating,
			Ott:         movie.MovieOtt,
			ReleaseDate: formatted_time,
			CategoryId:  uint32(movie.CategoryId),
			AdminId:     uint32(movie.AdminId),
		})
	}
	return movies,nil
}

func (db DBClient) UpdateProfile(req *proto.UpdateProfileRequest) (*User,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.Id).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	user.UserEmail = req.Email
	user.UserName = req.Name
	user.UserPhoneNumber = req.PhoneNumber
	if err := db.Db.Save(&user).Error; err != nil {
		return nil,status.Errorf(codes.Canceled, "Error while updating the user")
	}
	return &user,nil
}

func (db DBClient) AddReviewForMovie(req *proto.AddReviewRequest) (*Review,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	review := &Review{
		User_Id:     int(req.UserId),
		Movie_Id:    int(req.MovieId),
		Description: req.Description,
		Stars:       int(req.Stars),
	}

	if err := db.Db.Create(&review).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error while adding review for a movie: %v", err)
	}

	return review,nil
}

func (db DBClient) UpdateReviewForMovie(req *proto.UpdateReviewRequest) (*Review,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	var review Review
	result = db.Db.Model(&Review{}).Where("user_id=? and movie_id=?", req.UserId, req.MovieId).Find(&review)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil, status.Errorf(codes.NotFound, "User has not added review for this particular movie")
		}
	}

	review.Description = req.Description
	review.Stars = int(req.Stars)

	if err := db.Db.Save(&review).Error; err != nil {
		return nil, status.Errorf(codes.Canceled, "Error while saving a movie: %v", err)
	}

	return &review,nil
}

func (db DBClient) DeleteReviewForMovie(req *proto.DeleteReviewRequest) (uint32,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	var review Review
	result = db.Db.Model(&Review{}).Where("user_id=? and movie_id=?", req.UserId, req.MovieId).Find(&review)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "User has not added review to this particular movie")
		}
	}

	if err := db.Db.Unscoped().Delete(&review).Error; err != nil {
		return 500, status.Errorf(codes.Aborted, "Error while deleting a review for a movie")
	}

	return 200,nil
}

func (db DBClient) MarkAsRead(req *proto.MarkAsReadRequest) (*Viewed,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}
	viewed := &Viewed{
		User_Id:  int(req.UserId),
		Movie_Id: int(req.MovieId),
	}

	if err := db.Db.Create(&viewed).Error; err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Error while adding movie to Viewed Table")
	}

	return viewed,nil
}

func (db DBClient) MarkAsUnread(req *proto.MarkAsUnreadRequest) (uint32,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	var viewed Viewed
	result = db.Db.Model(&Viewed{}).Where("user_id=? and movie_id=?", req.UserId, req.MovieId).Find(&viewed)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "User has not added this particular movie into viewed")
		}
	}

	if err := db.Db.Unscoped().Delete(&viewed).Error; err != nil {
		return 500, status.Errorf(codes.Canceled, "Error while deleting record from Viewed Table")
	}

	return 200,nil
}

func (db DBClient) AddMovieToWatchList(req *proto.AddMovieToWatchListRequest) (*WatchList,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return nil,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}


	watchlist := &WatchList{
		User_Id: int(req.UserId),
		Movie_Id: int(req.MovieId),
	}

	if err := db.Db.Create(&watchlist).Error; err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Error while adding movie to watchlist: %v", err)
	}

	return watchlist,nil
}

func (db DBClient) RemoveMovieFromWatchList(req *proto.RemoveMovieFromWatchListRequest) (uint32,error){

	var user User
	result := db.Db.Model(&User{}).Where("id=?", req.UserId).Find(&user)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "User with provided Id doesn't exist in the User Table")
		}
	}

	var movie Movie
	result = db.Db.Model(&Movie{}).Where("id=?", req.MovieId).Find(&movie)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500,status.Errorf(codes.Canceled, "Movie with provided Id doesn't exist in the Movies Table")
		}
	}

	var watchlist WatchList
	result = db.Db.Model(&WatchList{}).Where("user_id=? and movie_id=?", req.UserId, req.MovieId).Find(&watchlist)
	if result.Error != nil {
		if result.RecordNotFound() == true {
			return 500, status.Errorf(codes.NotFound, "User has not added this particular movie into watchlist")
		}
	}

	if err := db.Db.Unscoped().Delete(&watchlist).Error; err != nil {
		return 500, status.Errorf(codes.Aborted, "Error while deleting movie from WatchList")
	}

	return 200,nil
}