package main

import "github.com/AJRDRGZ/go-db-gorm/storage"

func main() {
	driver := storage.MySQL
	storage.New(driver)

}
