package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/xryar/golang-grpc-ecommerce/internal/grpcmiddleware"
	"github.com/xryar/golang-grpc-ecommerce/internal/handler"
	"github.com/xryar/golang-grpc-ecommerce/internal/repository"
	"github.com/xryar/golang-grpc-ecommerce/internal/service"
	"github.com/xryar/golang-grpc-ecommerce/pb/auth"
	"github.com/xryar/golang-grpc-ecommerce/pb/product"
	"github.com/xryar/golang-grpc-ecommerce/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gocache "github.com/patrickmn/go-cache"
)

func main() {
	ctx := context.Background()
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	cacheService := gocache.New(time.Hour*24, time.Hour)

	authMiddleware := grpcmiddleware.NewAuthMiddleware(cacheService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, cacheService)
	authHandler := handler.NewAuthHandler(authService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware,
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(server, authHandler)
	product.RegisterProductServiceServer(server, productHandler)

	if os.Getenv("ENVIRONTMENT") == "dev" {
		reflection.Register(server)
		log.Println("Reflection is registered.")
	}

	log.Println("Server start in port :50051")
	if err := server.Serve(lis); err != nil {
		log.Panicf("Server error %v", err)
	}
}
