package dto

type CreateProductsStoreRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r *CreateProductsStoreRequest) FieldErrorCode(field string) string {
	switch field {
	case "name":
		return "ERR_PRODUCTS_STORE_NAME_REQUIRED"
	default:
		return "ERR"
	}
}
