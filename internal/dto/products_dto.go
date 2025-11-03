package dto

import "app/internal/domain"

type CreateProductsStoreRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateProductsStoreResponse struct {
	ProductsStore domain.ProductsStore `json:"products-store"`
}
