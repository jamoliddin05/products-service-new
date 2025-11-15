package helpers

import (
	"app/bootstrap/configs"
	"app/internal/handlers"
	"app/internal/middlewares"
	"app/internal/services"
	"app/internal/stores"
	"app/internal/uows"
	"app/internal/validators"
	"github.com/go-playground/validator/v10"
	"log"
)

func MustInitDB(cfg *configs.Config) *configs.Wrapper {
	dbWrapper, err := configs.NewDBWrapper(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	return dbWrapper
}

func BuildProductsStoreHandler(dbWrapper *configs.Wrapper) *handlers.GinProductsHandler {
	val := validators.NewValidator(validator.New())
	middleware := middlewares.NewRequestValidator(val)

	uow := uows.NewGormUnitOfWork[*stores.ProductsStoresOutboxStore](dbWrapper.DB(), stores.NewProductsStoresOutboxStore)
	productsStoresSvc := services.NewProductsService()
	outbox := services.NewOutboxService()
	productsHandler := handlers.NewGinProductsHandler(uow, middleware, productsStoresSvc, outbox)

	return productsHandler
}
