package dto

import "app/internal/domain"

type ProductsStoreResponse struct {
	ProductsStore *domain.ProductsStore `json:"products_store"`
}
