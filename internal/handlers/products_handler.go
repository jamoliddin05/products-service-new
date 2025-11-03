package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"app/internal/dto"
	"app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GinProductsHandler is the Gin adapter for UserService
type GinProductsHandler struct {
	productsService *services.ProductsService
}

// NewGinProductsHandler creates a new Gin adapter
func NewGinProductsHandler(productsService *services.ProductsService) *GinProductsHandler {
	return &GinProductsHandler{
		productsService: productsService,
	}
}

// BindRoutes registers the routes with Gin
func (h *GinProductsHandler) BindRoutes(r *gin.Engine) {
	auth := r.Group("/products")
	{
		auth.POST("/store/create", h.CreateStore)
	}
}

// CreateStore handles the /store/create route
func (h *GinProductsHandler) CreateStore(c *gin.Context) {
	var req dto.CreateProductsStoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorsMap := make(map[string]string)
			for _, fe := range ve {
				errorsMap[strings.ToLower(fe.Field())] = fe.Error()
			}
			JSONError(c, "Validation failed", http.StatusBadRequest, errorsMap)
			return
		}

		JSONError(c, err.Error(), http.StatusBadRequest, nil)
		return
	}

	userID := c.GetHeader("X-User-Id")
	userRoles := c.GetHeader("X-User-Roles")
	log.Printf("User roles header: %s", userRoles)
	resp, err := h.productsService.CreateProductsStore(req, userID, userRoles)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserNotSeller):
			JSONError(c, err.Error(), http.StatusForbidden, nil)
		case errors.Is(err, services.ErrInvalidCredentials):
			JSONError(c, err.Error(), http.StatusBadRequest, nil)
		default:
			JSONError(c, err.Error(), http.StatusInternalServerError, nil)
		}
		return
	}

	JSONSuccess(c, resp, http.StatusCreated)
}
