package server

import (
	"context"
	"log"
	"net"

	"example.com/pet-project/database"
	"example.com/pet-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func MockServer(mockDb *database.MockDatabase, ctx context.Context) (proto.MovieSuggestionsServiceClient, func()) {
	buffer := 1024 * 1024
	listen := bufconn.Listen(buffer)
	baseServer := grpc.NewServer()
	s := MoviesuggestionsServiceserver{
		Db: mockDb,
	}
	proto.RegisterMovieSuggestionsServiceServer(baseServer, &s)
	go func() {
		if err := baseServer.Serve(listen); err != nil {
			log.Printf("Error while serving the server")
		}
	}()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listen.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Error connecting to the server  %v", err)
	}

	closer := func() {
		err := listen.Close()
		if err != nil {
			log.Printf("Error while closing the server %v", err)
		}
		baseServer.Stop()
	}

	client := proto.NewMovieSuggestionsServiceClient(conn)

	return client, closer
}
