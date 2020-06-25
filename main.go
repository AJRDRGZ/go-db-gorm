package main

import (
	"fmt"

	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}
}
