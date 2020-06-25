package main

import (
	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)
}
