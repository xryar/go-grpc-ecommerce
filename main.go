package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/xryar/golang-grpc-ecommerce/internal/handler"
	"github.com/xryar/golang-grpc-ecommerce/pb/service"
	"github.com/xryar/golang-grpc-ecommerce/pkg/database"
	"github.com/xryar/golang-grpc-ecommerce/pkg/grpcmiddleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	serviceHandler := handler.NewServiceHandler()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
		),
	)

	service.RegisterHelloWorldServiceServer(server, serviceHandler)

	if os.Getenv("ENVIRONTMENT") == "dev" {
		reflection.Register(server)
		log.Println("Reflection is registered.")
	}

	log.Println("Server start in port :50051")
	if err := server.Serve(lis); err != nil {
		log.Panicf("Server error %v", err)
	}
}
