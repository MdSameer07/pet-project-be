package server

import (
	"database/sql"
	"time"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MoviesuggestionsServiceserver struct {
	proto.UnimplementedMovieSuggestionsServiceServer
	Db database.Database
	DB *gorm.DB
}

type Admin struct {
	gorm.Model
	AdminName     string `gorm:"not null"`
	AdminEmail    string `gorm:"not null;unique"`
	AdminPassword string `gorm:"not null"`
}

type User struct {
	gorm.Model
	UserName        string `gorm:"not null"`
	UserEmail       string `gorm:"not null;unique"`
	UserPassword    string `gorm:"not null"`
	UserPhoneNumber string `gorm:"not null;unique"`
}

type Movie struct {
	gorm.Model
	MovieName        string    `gorm:"not null"`
	MovieImage       string    `gorm:"not null;unique"`
	CategoryId       int       `gorm:"not null"`
	Category         Category  `gorm:"not null"`
	MovieRating      float32   `gorm:"not null"`
	MovieDirector    string    `gorm:"not null"`
	MovieDescription string    `gorm:"not null"`
	MovieReleaseDate time.Time `sql:"type:timestamp without time zone"`
	MovieOtt         string    `gorm:"not null"`
	AdminId          int       `gorm:"not null"`
}

type Category struct {
	gorm.Model
	Label string `gorm:"not null;unique"`
}

type WatchList struct {
	gorm.Model
	User_Id  int `gorm:"not null"`
	Movie_Id int `gorm:"not null"`
}

type Likes struct {
	gorm.Model
	User_Id  int `gorm:"not null"`
	Movie_Id int `gorm:"not null"`
}

type Viewed struct {
	gorm.Model
	User_Id  int `gorm:"not null"`
	Movie_Id int `gorm:"not null"`
}

type Review struct {
	gorm.Model
	User_Id     int    `gorm:"not null"`
	Movie_Id    int    `gorm:"not null"`
	Description string `gorm:"not null"`
	Stars       int    `gorm:"not null"`
}

type FeedBack struct {
	gorm.Model
	User_Id     int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (Viewed) TableName() string {
	return "viewed"
}

func (m *MoviesuggestionsServiceserver) AddingMoviesToSlice(rows *sql.Rows, movies []*proto.Movie) ([]*proto.Movie, error) {
	for rows.Next() {
		var movie Movie
		err := m.DB.ScanRows(rows, &movie)
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "Error while scanning movies into slice")
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
	return movies, nil
}