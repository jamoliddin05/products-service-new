package repositories

import (
	"app/internal/domain"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=ProductsStoreRepository --output=../mocks --structname=ProductsStoreRepositoryMock
type ProductsStoreRepository interface {
	Save(store *domain.ProductsStore) error
	GetStoreByUserId(id uuid.UUID) (*domain.ProductsStore, error)
}

type GormProductsStoreRepository struct {
	db *gorm.DB
}

func NewProductsStoreRepository(db *gorm.DB) ProductsStoreRepository {
	return &GormProductsStoreRepository{db: db}
}

func (r *GormProductsStoreRepository) Save(store *domain.ProductsStore) error {
	return r.db.Save(store).Error
}

func (r *GormProductsStoreRepository) GetStoreByUserId(id uuid.UUID) (*domain.ProductsStore, error) {
	var store domain.ProductsStore
	err := r.db.First(&store, "user_id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &store, nil
}
