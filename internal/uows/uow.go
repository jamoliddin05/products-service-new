package uows

import (
	"app/internal/repositories"
	"app/internal/stores"
	"gorm.io/gorm"
)

//go:generate mockery --name=UnitOfWork --output=../mocks --structname=UnitOfWorkMock
type UnitOfWork interface {
	Store() stores.Store
	DoCreateProductsStore(fn func(productsStoreRepo repositories.ProductsStoreRepository, eventRepo repositories.EventRepository) error) error
}

type gormUnitOfWork struct {
	db    *gorm.DB
	store stores.Store
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &gormUnitOfWork{
		db:    db,
		store: stores.NewStore(db),
	}
}

// Store plain store
func (u *gormUnitOfWork) Store() stores.Store {
	return u.store
}

// DoCreateProductsStore with tx
func (u *gormUnitOfWork) DoCreateProductsStore(fn func(repositories.ProductsStoreRepository, repositories.EventRepository) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		txStore := stores.NewStore(tx)
		return fn(txStore.ProductsStores(), txStore.Outbox())
	})
}
