package main

import (
	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	invoice := model.InvoiceHeader{
		Client: "Alvaro Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 3},
		},
	}

	storage.DB().Create(&invoice)
}
