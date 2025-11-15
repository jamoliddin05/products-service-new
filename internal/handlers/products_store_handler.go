package handlers

import (
	"app/internal/domain"
	"app/internal/middlewares"
	"app/internal/services"
	"app/internal/stores"
	"app/internal/uows"
	"errors"
	"net/http"

	"app/internal/dto"
	"github.com/gin-gonic/gin"
)

type GinProductsHandler struct {
	uow              uows.UnitOfWork[*stores.ProductsStoresOutboxStore]
	requestValidator *middlewares.RequestValidator
	productsStores   *services.ProductsStoreService
	outbox           *services.ProductsStoreOutboxService
}

func NewGinProductsHandler(
	uow uows.UnitOfWork[*stores.ProductsStoresOutboxStore],
	requestValidator *middlewares.RequestValidator,
	productsStores *services.ProductsStoreService,
	outbox *services.ProductsStoreOutboxService,
) *GinProductsHandler {
	return &GinProductsHandler{
		uow:              uow,
		requestValidator: requestValidator,
		productsStores:   productsStores,
		outbox:           outbox,
	}
}

func (h *GinProductsHandler) BindRoutes(r *gin.RouterGroup) {
	r.POST("/create", h.CreateStore)
	r.GET("/get", h.GetStoreByUserId)
}

func (h *GinProductsHandler) CreateStore(c *gin.Context) {
	var req dto.CreateProductsStoreRequest
	if !h.requestValidator.ValidateRequest(c, &req) {
		return
	}

	userId := c.GetHeader("X-User-Id")
	userRoles := c.GetHeader("X-User-Roles")
	var productsStore *domain.ProductsStore
	var err error
	err = h.uow.DoTransaction(func(store *stores.ProductsStoresOutboxStore) error {
		productsStore, err = h.productsStores.CreateProductsStore(store, req.Name, userId, userRoles)
		if err != nil {
			return err
		}

		err = h.outbox.SaveProductsStoreCreatedEvent(store, productsStore)
		if err != nil {
			return err
		}

		return nil
	})

	resp := dto.APIResponse{
		Errors: make(map[string]string),
	}
	status := http.StatusCreated

	if err != nil {
		switch {
		case errors.Is(err, services.ErrProductStoreAlreadyExists):
			resp.Errors["error"] = "ERR_PRODUCT_STORE_ALREADY_EXISTS"
			status = http.StatusConflict
		case errors.Is(err, services.ErrUserNotSeller):
			resp.Errors["error"] = "ERR_USER_NOT_SELLER"
			status = http.StatusForbidden
		case errors.Is(err, services.ErrInvalidCredentials):
			resp.Errors["error"] = "ERR_INVALID_CREDENTIALS"
			status = http.StatusBadRequest
		default:
			resp.Errors["error"] = "ERR_INTERNAL"
			status = http.StatusInternalServerError
		}
		c.JSON(status, resp)
		return
	}

	resp.Success = true
	resp.Data = dto.ProductsStoreResponse{
		ProductsStore: productsStore,
	}

	c.JSON(status, resp)
}

func (h *GinProductsHandler) GetStoreByUserId(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	userRoles := c.GetHeader("X-User-Roles")
	var productsStore *domain.ProductsStore
	var err error
	err = h.uow.DoTransaction(func(store *stores.ProductsStoresOutboxStore) error {
		productsStore, err = h.productsStores.GetProductsStoreByUserID(store, userId, userRoles)
		if err != nil {
			return err
		}

		return nil
	})

	resp := dto.APIResponse{
		Errors: make(map[string]string),
	}
	status := http.StatusOK

	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserNotSeller):
			resp.Errors["error"] = "ERR_USER_NOT_SELLER"
			status = http.StatusForbidden
		case errors.Is(err, services.ErrInvalidCredentials):
			resp.Errors["error"] = "ERR_INVALID_CREDENTIALS"
			status = http.StatusBadRequest
		default:
			resp.Errors["error"] = "ERR_INTERNAL"
			status = http.StatusInternalServerError
		}
		c.JSON(status, resp)
		return
	}

	resp.Success = true
	resp.Data = dto.ProductsStoreResponse{
		ProductsStore: productsStore,
	}

	c.JSON(status, resp)
}
