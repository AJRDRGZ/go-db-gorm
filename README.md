## Migraciones

```go
storage.DB().AutoMigrate(
    &model.Product{},
    &model.InvoiceHeader{},
    &model.InvoiceItem{},
)
```

## Crear Productos

```go
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
```
