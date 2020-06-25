package main

import (
	"fmt"

	"github.com/AJRDRGZ/go-db-gorm/model"
	"github.com/AJRDRGZ/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	myProduct := model.Product{}

	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)
}
