package data

import (
	"log/slog"
	"os"

	"globalco/productAPI/models"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {

	logger.Info("Creating new ProductRepository")

	return &ProductRepository{}
}

func (pr *ProductRepository) GetAllProducts() []models.Product {
	// Placeholder implementation
	logger.Info("Fetching all products from repository")
	return []models.Product{}
}
