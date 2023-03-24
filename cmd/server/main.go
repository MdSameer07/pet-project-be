package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"example.com/pet-project/database"
	"example.com/pet-project/gen/proto"
	"example.com/pet-project/server"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)


var port = flag.Int("port", 50059, "The server port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen to the server: %v", err)
		return
	}

	err = godotenv.Load("/home/mohammed/Documents/Full-Stack/pet-project/pet-project-be/.env")
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

	server1 := grpc.NewServer()
	proto.RegisterMovieSuggestionsServiceServer(server1, &server.MoviesuggestionsServiceserver{
		Db: database.DBClient{Db: db},
		DB: db,
	})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server1.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}
}
