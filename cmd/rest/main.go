package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/xryar/golang-grpc-ecommerce/internal/handler"
	"github.com/xryar/golang-grpc-ecommerce/internal/repository"
	"github.com/xryar/golang-grpc-ecommerce/internal/service"
	"github.com/xryar/golang-grpc-ecommerce/pkg/database"
)

func handleGetFileName(c *fiber.Ctx) error {
	fileNameParam := c.Params("filename")
	filePath := path.Join("storage", "product", fileNameParam)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(http.StatusNotFound).SendString("Not Found")
		}

		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	ext := path.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	c.Set("Content-Type", mimeType)
	return c.SendStream(file)
}

func main() {
	godotenv.Load()
	ctx := context.Background()
	app := fiber.New()

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Connected to database")

	orderRepository := repository.NewOrderRepository(db)
	webhookService := service.NewWebhookService(orderRepository)
	webhookHandler := handler.NewWebhookHandler(webhookService)

	app.Use(cors.New())
	app.Get("/storage/product/:filename", handleGetFileName)
	app.Post("/product/upload", handler.UploadProductImageHandler)
	app.Post("/webhook/xendit/invoice", webhookHandler.ReceiveInvoice)

	app.Listen(":3000")
}
