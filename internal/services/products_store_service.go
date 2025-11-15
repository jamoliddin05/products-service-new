package services

import (
	"app/internal/domain"
	"app/internal/stores"
	"errors"
	"github.com/google/uuid"
	"strings"
)

var (
	ErrUserNotSeller             = errors.New("user isn't a seller")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrProductStoreAlreadyExists = errors.New("product store already exists")
	ProductsStoreCreated         = "ProductsStoreCreated"
)

type ProductsStoreService struct{}

func NewProductsService() *ProductsStoreService {
	return &ProductsStoreService{}
}

func (s *ProductsStoreService) CreateProductsStore(
	store *stores.ProductsStoresOutboxStore,
	storeName string,
	userId string,
	roles string,
) (*domain.ProductsStore, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !isSeller(roles) {
		return nil, ErrUserNotSeller
	}

	productsStore, err := store.ProductsStores().GetStoreByUserId(userUUID)
	if err != nil {
		return nil, err
	}
	if productsStore != nil {
		return nil, ErrProductStoreAlreadyExists
	}

	productsStore = &domain.ProductsStore{
		UserID: userUUID,
		Name:   storeName,
	}

	if err := store.ProductsStores().Save(productsStore); err != nil {
		return nil, err
	}

	return productsStore, nil
}

func (s *ProductsStoreService) GetProductsStoreByUserID(
	store *stores.ProductsStoresOutboxStore,
	userId string,
	roles string,
) (*domain.ProductsStore, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !isSeller(roles) {
		return nil, ErrUserNotSeller
	}

	productsStore, err := store.ProductsStores().GetStoreByUserId(userUUID)
	if err != nil {
		return nil, err
	}

	return productsStore, nil
}

func isSeller(roles string) bool {
	var userRoles []string
	if roles != "" {
		userRoles = strings.Split(roles, ",")
		for i := range userRoles {
			userRoles[i] = strings.TrimSpace(userRoles[i])
		}
	}

	isSeller := false
	for _, r := range userRoles {
		if strings.EqualFold(r, "seller") {
			isSeller = true
			break
		}
	}

	return isSeller
}
