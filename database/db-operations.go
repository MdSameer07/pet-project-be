package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

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

type Category struct{
 	gorm.Model
	Label string       `gorm:"not null;unique"`
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

func (Viewed) TableName() string{
	return "viewed"
}

func main() {
	err := godotenv.Load("/home/mohammed/Documents/Full-Stack/pet-project/pet-project-be/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}
	Db_details := os.Getenv("DB_Details")

	db, err := gorm.Open("postgres", Db_details)
	if err != nil {
		log.Fatalf("Connection with database failed: %v", err)
		return
	}
	defer db.Close()
	dbase := db.DB()
	defer dbase.Close()
	err = dbase.Ping()
	if err != nil {
		log.Fatalf("Pinging database Failed: %v", err)
		return
	}
	fmt.Println("Connection to database was successfully established")

	if err := db.AutoMigrate(&Admin{}, &User{}, &Movie{}, &WatchList{}, &Likes{}, &Viewed{}, &FeedBack{},&Category{},&Review{}).Error; err != nil {
		log.Fatalf("Error while automigrating tables: %v", err)
		return
	}
	fmt.Println("Success in creating tables")

	if err := db.Model(&Movie{}).AddForeignKey("admin_id", "admins(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Movie table: %v", err)
		return
	}
	if err := db.Model(&Movie{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Movie table: %v", err)
		return
	}
	if err := db.Model(&WatchList{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to WatchList table: %v", err)
		return
	}
	if err := db.Model(&WatchList{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to WatchList table: %v", err)
		return
	}
	if err := db.Model(&Likes{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Likes table: %v", err)
		return
	}
	if err := db.Model(&Likes{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Likes table: %v", err)
		return
	}
	if err := db.Model(&Viewed{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Insight table: %v", err)
		return
	}
	if err := db.Model(&Viewed{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Insight table: %v", err)
		return
	}
	if err := db.Model(&FeedBack{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Feedback table: %v", err)
		return
	}
	if err := db.Model(&Review{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Review table: %v", err)
		return
	}
	if err := db.Model(&Review{}).AddForeignKey("movie_id", "movies(id)", "CASCADE", "CASCADE").Error; err != nil {
		log.Fatalf("Error while adding foreign key to Review table: %v", err)
		return
	}

	fmt.Println("Success adding foreign keys")

	if err := db.Table("watch_lists").AddUniqueIndex("idx_wusers_wmovies", "user_id", "movie_id").Error; err!=nil{
		log.Fatalf("Error while adding unique indexes to watchlist table: %v",err)
		return
	}

	if err := db.Table("likes").AddUniqueIndex("idx_lusers_lmovies", "user_id", "movie_id").Error; err!=nil{
		log.Fatalf("Error while adding unique indexes to likes table: %v",err)
		return
	}

	if err := db.Table("viewed").AddUniqueIndex("idx_vusers_vmovies", "user_id", "movie_id").Error; err!=nil{
		log.Fatalf("Error while adding unique indexes to viewed table: %v",err)
		return
	}

	if err := db.Table("reviews").AddUniqueIndex("idx_rusers_rmovies", "user_id", "movie_id").Error; err!=nil{
		log.Fatalf("Error while adding unique indexes to review table: %v",err)
		return
	}

	if err := db.Table("feed_backs").AddUniqueIndex("idx_user_desc", "user_id", "description").Error; err!=nil{
		log.Fatalf("Error while adding unique indexes to review table: %v",err)
		return
	}

	fmt.Println("Success adding unique indexes to tables")

	Admin1_Password := os.Getenv("Admin1_Password")

	admin1 := Admin{
		AdminName:     "Sameer",
		AdminEmail:    "md.sameer@beautifulcode.in",
		AdminPassword: Admin1_Password,
	}

	if err := db.Create(&admin1).Error; err != nil {
		log.Fatalf("Error while creating Admin1: %v", err)
	}

	fmt.Println("Admin1 created successfully")

	Admin2_Password := os.Getenv("Admin2_Password")

	admin2 := Admin{
		AdminName:     "Tarun",
		AdminEmail:    "taruntrs@gmail.com",
		AdminPassword: Admin2_Password,
	}

	if err := db.Create(&admin2).Error; err != nil {
		log.Fatalf("Error while creating Admin2: %v", err)
	}

	fmt.Println("Admin2 created successfully")

	movie1 := Movie{
		MovieName: "The Dark Knight",
		MovieImage:  "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/06/The-Joker-from-The-Dark-Knight.jpg?q=50&fit=crop&w=1500&dpr=1.5",
		MovieRating: 9.0,
		MovieDirector: "Cristopher Nolan",
		MovieDescription: "The plot follows the vigilante Batman, police lieutenant James Gordon, and district attorney Harvey Dent, who form an alliance to dismantle organized crime in Gotham City. Their efforts are derailed by the Joker, an anarchistic mastermind who seeks to test how far Batman will go to save the city from chaos.",
		MovieReleaseDate: time.Date(2008,7,18,0,0,0,0,time.UTC),
		MovieOtt: "https://www.primevideo.com/",
		Category: Category{
			Label: "thriller",
		},
		AdminId: 1,
	}

	if err := db.Create(&movie1).Error; err!=nil{
		log.Fatalf("Error while creating movie1: %v",err)
		return 
	}

	fmt.Println("Movie1 created successfully")

	movie2 := Movie{
		MovieName: "The Exorcist",
		MovieImage:  "https://prd-rteditorial.s3.us-west-2.amazonaws.com/wp-content/uploads/2020/10/28144852/Scariest_Movies_Exorcist.jpg",
		MovieRating: 8.1,
		MovieDirector: "John Boorman",
		MovieDescription: "When a teenage girl is possessed by a mysterious entity, her mother seeks the help of two priests to save her daughter. A visiting actress in Washington, D.C., notices dramatic and dangerous changes in the behavior and physical make-up of her 12-year-old daughter.",
		MovieReleaseDate: time.Date(1973,12,26,0,0,0,0,time.UTC),
		MovieOtt: "https://www.hulu.com/",
		Category: Category{
			Label: "horror",
		},
		AdminId: 1,
	}

	if err := db.Create(&movie2).Error; err!=nil{
		log.Fatalf("Error while creating movie2: %v",err)
		return 
	}

	fmt.Println("Movie2 created successfully")

	movie3 := Movie{
		MovieName: "The Lord Of The Rings",
		MovieImage:  "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/10/returnoftheking.jpg?q=50&fit=crop&w=1500&dpr=1.5",
		MovieRating: 9.0,
		MovieDirector:  "Peter Jackson",
		MovieDescription: "The final confrontation between the forces of good and evil fighting for control of the future of Middle-earth. Frodo and Sam reach Mordor in their quest to destroy the One Ring, while Aragorn leads the forces of good against Sauron's evil army at the stone city of Minas Tirith.",
		MovieReleaseDate: time.Date(2004,2,6,0,0,0,0,time.UTC),
		MovieOtt: "https://www.netflix.com/in/",
		Category: Category{
			Label: "action",
		},
		AdminId: 1,
	}

	if err := db.Create(&movie3).Error; err!=nil{
		log.Fatalf("Error while creating movie3: %v",err)
		return 
	}

	fmt.Println("Movie3 created successfully")

	movie4 := Movie{
		MovieName: "Home Alone",
		MovieImage:  "https://www.dkoding.in/wp-content/uploads/Home-Alone-Part-6-Release-Date-Hollywood-Entertainment-DKODING.jpg",
		MovieRating: 7.7,
		MovieDirector:  "Chris Columbus",
		MovieDescription: "When bratty 8-year-old Kevin McCallister (Macaulay Culkin) acts out the night before a family trip to Paris, his mother (Catherine O'Hara) makes him sleep in the attic. After the McCallisters mistakenly leave for the airport without Kevin, he awakens to an empty house and assumes his wish to have no family has come true. But his excitement sours when he realizes that two con men (Joe Pesci, Daniel Stern) plan to rob the McCallister residence, and that he alone must protect the family home.",
		MovieReleaseDate: time.Date(1991,10,18,0,0,0,0,time.UTC),
		MovieOtt: "https://www.hotstar.com/in",
		Category: Category{
			Label: "comedy",
		},
		AdminId: 2,
	}

	if err := db.Create(&movie4).Error; err!=nil{
		log.Fatalf("Error while creating movie3: %v",err)
		return 
	}

	fmt.Println("Movie4 created successfully")
}
