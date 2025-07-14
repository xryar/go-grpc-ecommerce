package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xryar/golang-grpc-ecommerce/internal/handler"
)

func main() {
	app := fiber.New()

	app.Post("/product/upload", handler.UploadProductImageHandler)

	app.Listen(":3000")
}
