package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/xryar/golang-grpc-ecommerce/internal/handler"
	"github.com/xryar/golang-grpc-ecommerce/pb/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	serviceHandler := handler.NewServiceHandler()

	server := grpc.NewServer()

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
