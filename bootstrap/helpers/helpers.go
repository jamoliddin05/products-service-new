package helpers

import (
	"app/bootstrap/configs"
	"app/internal/handlers"
	"app/internal/services"
	"app/internal/uows"
	"log"
)

// MustInitDB инициализирует базу данных или падает
func MustInitDB(cfg *configs.Config) *configs.Wrapper {
	dbWrapper, err := configs.NewDBWrapper(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	return dbWrapper
}

// BuildProductsHandler строит слой ProductsService + Gin handler
func BuildProductsHandler(dbWrapper *configs.Wrapper) *handlers.GinProductsHandler {
	uow := uows.NewUnitOfWork(dbWrapper.DB())
	productsSvc := services.NewProductsService(uow)
	productsHandler := handlers.NewGinProductsHandler(productsSvc)
	return productsHandler
}
