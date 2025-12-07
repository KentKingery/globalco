package main

import (
	"log/slog"
	"os"

	"github.com/google/uuid"

	"globalco/productAPI/models"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	product := models.Product{
		ID:    uuid.New(),
		Name:  "Sample Product",
		Price: 29.99,
	}
	logger.Info("Product API service started")
	logger.Info("Created product", "product", product)
	logger.Info("Product API service started")
}
