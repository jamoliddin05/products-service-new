package repositories

import (
	"app/internal/domain"
	"gorm.io/gorm"
)

//go:generate mockery --name=ProductsStoreRepository --output=../mocks --structname=ProductsStoreRepositoryMock
type ProductsStoreRepository interface {
	Create(store *domain.ProductsStore) error
}

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsStoreRepository(db *gorm.DB) ProductsStoreRepository {
	return &StoreRepositoryImpl{db: db}
}

func (r *StoreRepositoryImpl) Create(store *domain.ProductsStore) error {
	return r.db.Create(store).Error
}
