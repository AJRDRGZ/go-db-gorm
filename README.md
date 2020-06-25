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

## Consultar todos los productos

```go
products := make([]model.Product, 0)
storage.DB().Find(&products)

for _, product := range products {
    fmt.Printf("%d - %s\n", product.ID, product.Name)
}
```

## Consultar un solo producto

```go
myProduct := model.Product{}

storage.DB().First(&myProduct, 3)
fmt.Println(myProduct)
```

# Actualizar un Producto

```go
myProduct := model.Product{}
myProduct.ID = 3

storage.DB().Model(&myProduct).Updates(
    model.Product{Name: "Curso de Java", Price: 120},
)
```
