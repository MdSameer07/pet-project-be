package database

import (
	"database/sql/driver"
	"testing"
	"time"

	"example.com/pet-project/gen/proto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func newMockDBClient() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic("Got an unexpected error while opening mock sql connection.")
	}
	return gormDB, mock
}

func TestDeleteMovieFromDatabase(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	DBClient := DBClient{
		Db: db,
	}

	result := sqlmock.NewResult(7, 1)
	t.Run("Deleting Movie", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "movies".*$`).WillReturnRows(sqlmock.NewRows([]string{"id", "movie_name", "movie_image", "category_id", "movie_rating", "movie_director", "movie_description", "movie_release_date", "movie_ott", "admin_id"}).AddRow(7, 'a', 'b', 14, 7.7, 'c', 'd', time.Date(2008, 7, 18, 0, 0, 0, 0, time.UTC), 'f', 21))
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "movies" WHERE "movies"."id" = ?`).WithArgs(7).WillReturnResult(result)
		mock.ExpectCommit()

		_, err := DBClient.DeleteMovieFromDatabase(&proto.DeleteMovieFromDatabaseRequest{MovieId: 7})
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}
		assert.Nil(t, err)
	})
}

func TestAddMovieToDatabase(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	DBClient := DBClient{
		Db: db,
	}

	t.Run("Adding Movie", func(t *testing.T) {
		// result := sqlmock.NewResult(1, 1)
		releaseDate, _ := time.Parse("02-01-2006", "18-07-2008")
		mock.ExpectQuery(`SELECT (.+) FROM "admins".*$`).WithArgs(5).WillReturnRows(sqlmock.NewRows([]string{"id", "admin_name", "admin_email"}).AddRow(5, "Rohith", "rohith@gmail.com"))
		mock.ExpectQuery(`SELECT (.+) FROM "categories".*$`).WithArgs("thriller").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(7))

		// mock.ExpectBegin()
		// mock.ExpectExec(`INSERT INTO "categories".*$`).WithArgs(time.Now(),time.Now(),sql.NullTime{},"thriller").WillReturnResult(sqlmock.NewResult(7,1))
		// mock.ExpectCommit()
		mock.ExpectBegin()
		// mock.ExpectExec(`INSERT INTO "movies" ("created_at","updated_at","deleted_at","movie_name","movie_image","category_id","movie_rating","movie_director","movie_description","movie_release_date","movie_ott","admin_id")VALUES (?,?,?,?,?,?,?,?,?,?,?)`).WithArgs(time.Now(),time.Now(),nil,"a","b",7,7.7,"c","d",time.Date(2008, 7, 18, 0, 0, 0, 0, time.UTC),"e",5).WillReturnResult(result)
		mock.ExpectQuery(`INSERT INTO "movies" (.+)`).WithArgs(sqlmock.AnyArg(),sqlmock.AnyArg(),sqlmock.AnyArg(),"a","b",7,float32(7.7),"c","d",releaseDate,"e",5,
).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
	})

	_, err := DBClient.AddMovieToDatabase(&proto.AddMovieToDatabaseRequest{Name: "a", Imageurl: "b", Category: "Thriller", Rating: 7.7, Director: "c", Description: "d", ReleaseDate: "18-07-2008", Movieott: "e", AdminId: 5})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
	assert.Nil(t, err)
}
