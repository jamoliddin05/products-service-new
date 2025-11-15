package stores

import (
	"app/internal/repositories"
	"gorm.io/gorm"
)

type ProductsStoresOutboxStore struct {
	db *gorm.DB
}

func NewProductsStoresOutboxStore(db *gorm.DB) *ProductsStoresOutboxStore {
	return &ProductsStoresOutboxStore{db: db}
}

func (s *ProductsStoresOutboxStore) ProductsStores() repositories.ProductsStoreRepository {
	return repositories.NewProductsStoreRepository(s.db)
}

func (s *ProductsStoresOutboxStore) Outbox() repositories.EventRepository {
	return repositories.NewEventRepository(s.db)
}
