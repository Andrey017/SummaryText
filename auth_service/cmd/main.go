package main

import (
	"auth_service/pkg/config"
	"auth_service/pkg/db"
	"auth_service/pkg/pb"
	"auth_service/pkg/services"
	"auth_service/pkg/utils"
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

	serverDb := db.Init(c.DBUrl)

	jwt := utils.JWTWrapper{
		SecretKey:       c.JwtSecretKey,
		Issuer:          "auth_service",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Faied to listing", err)
	}

	fmt.Println("Auth service on", c.Port)

	s := services.Server{
		ServerDB: serverDb,
		Jwt:      jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
