package main

import (
	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)
}
