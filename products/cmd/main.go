package main

import (
	"github.com/juancanchi/jujuy-market/products/internal/infrastructure/postgres"
	"github.com/juancanchi/jujuy-market/products/internal/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	handler "github.com/juancanchi/jujuy-market/products/internal/delivery/http"
)

func main() {
	db, err := postgres.NewDB("host=localhost user=postgres password=postgres dbname=jujuy_market port=5432 sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	repo := postgres.NewProductRepository(db)
	uc := usecase.NewProductUsecase(repo)
	h := handler.NewProductHandler(uc)

	r := gin.Default()

	// Rutas
	r.POST("/products", h.Create)
	r.GET("/products", h.List)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
