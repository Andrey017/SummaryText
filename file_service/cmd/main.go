package main

import (
	"file_service/pkg/client"
	"file_service/pkg/config"
	"file_service/pkg/pb"
	"file_service/pkg/services"
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

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Println("Article service on", c.Port)

	articleSvc := client.InitArticleServiceClient(c.ArticleSvc)

	s := services.Server{
		ArticleClient: articleSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFileServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
