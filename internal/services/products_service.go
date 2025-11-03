package services

import (
	"app/internal/domain"
	"app/internal/dto"
	"app/internal/uows"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var (
	ErrUserNotSeller      = errors.New("user isn't a seller")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type ProductsService struct {
	uow uows.UnitOfWork
}

func NewProductsService(uow uows.UnitOfWork) *ProductsService {
	return &ProductsService{uow: uow}
}

func (s *ProductsService) CreateProductsStore(req dto.CreateProductsStoreRequest, userId string, roles string) (*dto.CreateProductsStoreResponse, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

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

	if !isSeller {
		return nil, ErrUserNotSeller
	}

	productsStore := domain.ProductsStore{
		UserID: userUUID,
		Name:   req.Name,
	}

	if err := s.uow.Store().ProductsStores().Create(&productsStore); err != nil {
		return nil, fmt.Errorf("failed to create products store: %w", err)
	}

	return &dto.CreateProductsStoreResponse{
		ProductsStore: productsStore,
	}, nil
}
