package main

import (
	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Delete Soft
	myProduct := model.Product{}
	myProduct.ID = 4

	storage.DB().Delete(&myProduct)

	// Delete Permanent
	myProduct = model.Product{}
	myProduct.ID = 4

	storage.DB().Unscoped().Delete(&myProduct)
}
