package main

import (
	"article_service/pkg/config"
	"article_service/pkg/db"
	"article_service/pkg/pb"
	"article_service/pkg/services"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	serverDB := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Println("Article service on", c.Port)

	s := services.Server{
		ServerDB: serverDB,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
