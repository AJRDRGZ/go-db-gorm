package main

import (
	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	product1 := model.Product{
		Name:  "Curso de Go",
		Price: 120,
	}

	obs := "testing with Go"
	product2 := model.Product{
		Name:         "Curso de Testing",
		Price:        150,
		Observations: &obs,
	}

	product3 := model.Product{
		Name:  "Curso de Python",
		Price: 200,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
}
