package main

import (
	"context"
	"io"
	"log"

	"example.com/pet-project/proto"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return
	}
	defer connection.Close()

	client := proto.NewMovieSuggestionsServiceClient(connection)

	// Adding Movie To Database

	resp1, err := client.AddMovieToDatabase(context.Background(), &proto.AddMovieToDatabaseRequest{Name: "Titanic", Category: "Romance", Rating: 7.9, Director: "James Cameron", Description: "James Cameron's 'Titanic' is an epic, action-packed romance set against the ill-fated maiden voyage of the R.M.S. Titanic; the pride and joy of the White Star Line and, at the time, the largest moving object ever built. She was the most luxurious liner of her era -- the 'ship of dreams' -- which ultimately carried over 1,500 people to their death in the ice cold waters of the North Atlantic in the early hours of April 15, 1912.", ReleaseDate: "19-10-1997", Movieott: "https://www.hotstar.com/in", Imageurl: "https://deadline.com/wp-content/uploads/2021/02/MCDTITA_FE014.jpg", AdminId: 2})
	if err != nil {
		log.Fatalf("Failed to add movie to the database: %v", err)
		return
	}
	log.Printf("Movie added Successfully: %v", resp1.Movie)

	// deleting Movie from Database

	resp2, err := client.DeleteMovieFromDatabase(context.Background(), &proto.DeleteMovieFromDatabaseRequest{MovieId: 4})
	if err != nil {
		log.Println(resp2.Error)
		return
	}
	log.Printf("Movie deleted successfully with status: %v", resp2.Status)

	//Getting feedback

	resp3, err := client.GetFeedBack(context.Background(), &proto.GetFeedBackRequest{AdminEmail: "md.sameer@beautifulcode.in"})
	if err != nil {
		log.Fatalf("Error while getting feedbacks from database: %v", err)
		return
	}
	for {
		resp, err := resp3.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading a feedback: %v", err)
			return
		}
		log.Println(resp.Description)
	}

	//User Login

	resp4, err := client.Login(context.Background(), &proto.LoginRequest{Email: "shreekar69@gmail.com", Password: "Shreekar@123"})
	if err != nil {
		log.Fatalf("Error while logging in User: %v", err)
		return
	}
	log.Printf("User login successfully : %v", resp4)

	//User Register

	resp5, err := client.Register(context.Background(), &proto.RegisterRequest{Name: "Shreekar", Email: "shreekar69@gmail.com", Password: "Shreekar@123", PhoneNumber: "9384723839"})
	if err != nil {
		log.Fatalf("Error while registering a user : %v", err)
		return
	}
	log.Printf("User registered Successfully: %v", resp5)

	//Getting All Movies From Database

	resp6, err := client.GetAllMovies(context.Background(), &proto.GetAllMoviesRequest{})
	if err != nil {
		log.Fatalf("Error while getting movies from database: %v", err)
		return
	}
	for {
		resp, err := resp6.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(err.Error())
			log.Fatalf("Error while getting a movie: %v", err)
			return
		}
		log.Println(resp.Movie)
	}

	//Searching movies based on name or category

	resp7, err := client.SearchForMovies(context.Background(), &proto.SearchRequest{
		Filter: proto.SearchRequest_Category,
		Category: &proto.Category{
			Type: proto.Category_Comedy,
		},
	})

	// resp7,err := client.SearchForMovies(context.Background(),&proto.SearchRequest{
	// 	Filter: proto.SearchRequest_Name,
	// 	Name: "the",
	// })

	if err != nil {
		log.Fatalf("Error getting movies from database: %v", err)
		return
	}
	for {
		resp, err := resp7.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while getting a movie: %v", err)
			return
		}
		log.Println(resp.Movie)
	}

	//Adding movie to WatchList

	resp8, err := client.AddMovieToWatchList(context.Background(), &proto.AddMovieToWatchListRequest{UserId: 1, MovieId: 1})
	if err != nil {
		log.Fatalf("Error while adding movie to watchlist: %v", err)
		return
	}
	log.Println(resp8)

	//Deleting Movie from WatchList

	resp9, err := client.RemoveMovieFromWatchList(context.Background(), &proto.RemoveMovieFromWatchListRequest{UserId: 1, MovieId: 3})
	if err != nil {
		log.Fatalf(resp9.Errors)
		return
	}
	log.Printf("Successfully deleted Movie from Wathclist with status : %v", resp9.Status)

	//Adding Movie to Likes

	resp10, err := client.AddMovieToLikes(context.Background(), &proto.AddMovieToLikesRequest{UserId: 3, MovieId: 1})
	if err != nil {
		log.Fatalf("Error while adding movie to likes: %v", err)
		return
	}
	log.Println(resp10)

	//Deleting Movie From Likes

	resp11, err := client.RemoveMovieFromLikes(context.Background(), &proto.RemoveMovieFromLikesRequest{UserId: 2, MovieId: 4})
	if err != nil {
		log.Println(resp11.Errors)
		return
	}
	log.Printf("Movie deleted successfully with status: %v", resp11.Status)

	//Adding review for a movie

	resp12, err := client.AddReviewForMovie(context.Background(), &proto.AddReviewRequest{UserId: 2, MovieId: 6, Description: "It's been several decades since I had the chance to watch James Cameron's masterful epic, and watching it again after all this time, I can honestly say that my feelings haven't changed: this is an exhilarating film from top to bottom, from the magnificent visuals and a love story that deserves its place among Hollywood's best.", Stars: 4})
	if err != nil {
		log.Fatalf("Error while adding a review: %v", err)
		return
	}
	log.Println(resp12)

	//Update Review for a movie

	resp13, err := client.UpdateReviewForMovie(context.Background(), &proto.UpdateReviewRequest{UserId: 2, MovieId: 3, Stars: 4, Description: "The Fellowship of the Ring: Not just a Movie, but the Door to another Dimension.The first part of the Lord of the Rings trilogy, the Fellowship of the Rings opened the door to a whole new world for me. I'd never read any of Tolkien's books when I saw the film for the first time at the theatre and, now that I've read them, in retrospect I think being a neophyte to the mythology made my LOTR movie experience all the more miraculous."})
	if err != nil {
		log.Fatalf("Error while updating a review: %v", err)
		return
	}
	log.Println(resp13)

	//Delete Review for a movie

	resp14, err := client.DeleteReviewForMovie(context.Background(), &proto.DeleteReviewRequest{UserId: 2, MovieId: 6})
	if err != nil {
		log.Fatalf(resp14.Errors)
		return
	}
	log.Printf("Review Deleted Successfully with status : %v", resp14.Status)

	//Updating Profile

	resp15, err := client.UpdateProfile(context.Background(), &proto.UpdateProfileRequest{Id: 1, Name: "Shreekar", Email: "shreekar69@gmail.com", PhoneNumber: "9283746519"})
	if err != nil {
		log.Fatalf("Error while updating profile: %v", err)
		return
	}
	log.Println(resp15)

	//Marking movies as watched

	resp16, err := client.MarkAsRead(context.Background(), &proto.MarkAsReadRequest{UserId: 1, MovieId: 3})
	if err != nil {
		log.Fatalf("Error while adding movies to viewed Table: %v", err)
		return
	}
	log.Println(resp16)

	//Unmarking movies as watched

	resp17, err := client.MarkAsUnread(context.Background(), &proto.MarkAsUnreadRequest{UserId: 1, MovieId: 1})
	if err != nil {
		log.Fatalf(resp17.Errors)
		return
	}
	log.Printf("Successfully deleted movie from viewed Table with status : %v", resp17.Status)

	//Adding a Feedback

	resp18, err := client.GiveFeedBack(context.Background(), &proto.GiveFeedBackRequest{UserId: 2, Description: "Great App for finding perfect movies!! wanna insist to add more movies"})
	if err != nil {
		log.Fatalf("Error while adding a feedback: %v", err)
		return
	}
	log.Println(resp18)

	//Updating a feedback

	resp19, err := client.UpdateFeedBack(context.Background(), &proto.UpdateFeedBackRequest{UserId: 3, FeedbackId: 2, Description: "Imporved a bit. A little more would be great."})
	if err != nil {
		log.Fatalf("Error while updating a feedback: %v", err)
		return
	}
	log.Println(resp19)

	//Deleting a feedback

	resp20, err := client.DeleteFeedBack(context.Background(), &proto.DeleteFeedBackRequest{UserId: 3, FeedbackId: 2})
	if err != nil {
		log.Fatalf(resp20.Errors)
		return
	}
	log.Printf("Successfully deleted feedback from FeedBacks table with status : %v", resp20.Status)
}
