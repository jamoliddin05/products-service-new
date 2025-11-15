package main

import (
	"app/bootstrap"
	"app/bootstrap/configs"
	"app/bootstrap/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()
	dbWrapper := helpers.MustInitDB(cfg)
	productsStoreHandler := helpers.BuildProductsStoreHandler(dbWrapper)

	r := gin.Default()
	storeGroup := r.Group("/products/store")
	productsStoreHandler.BindRoutes(storeGroup)

	app := bootstrap.NewApp(r, ":8080")
	app.RegisterCloser(dbWrapper)

	app.RunWithGracefulShutdown()
}
