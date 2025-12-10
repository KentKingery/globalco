package main

import (
	"log/slog"
	"os"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"

	"globalco/productAPI/data"
	"globalco/productAPI/models"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {

	opt := badger.DefaultOptions("")
	opt.Dir = "./tmp/badger"
	opt.ValueDir = "./tmp/badger"

	db, err := badger.Open(opt)
	if err != nil {
		slog.Error(err.Error())
	}

	defer db.Close()

	product := models.Product{
		ID:       uuid.New(),
		Name:     "Sample Product",
		Price:    29.99,
		Quantity: 10,
	}

	logger.Info("Product API service started")
	logger.Info("Created product", "product", product)

	repo := data.NewProductRepository()
	repo.GetAllProducts()
	logger.Info("Product API service stopped")
}
