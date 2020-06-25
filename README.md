## Migraciones

```go
storage.DB().AutoMigrate(
    &model.Product{},
    &model.InvoiceHeader{},
    &model.InvoiceItem{},
)
```
