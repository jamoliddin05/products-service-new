package services

import (
	"app/internal/stores"
)

type ProductsStoreOutboxService struct {
}

func NewOutboxService() *ProductsStoreOutboxService {
	return &ProductsStoreOutboxService{}
}

func (s *ProductsStoreOutboxService) SaveProductsStoreCreatedEvent(store *stores.ProductsStoresOutboxStore, payload interface{}) error {
	return store.Outbox().Save(ProductsStoreCreated, payload)
}
