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
	productsHandler := helpers.BuildProductsHandler(dbWrapper)
	r := gin.Default()
	productsHandler.BindRoutes(r)

	app := bootstrap.NewApp(r, ":8080")
	app.RegisterCloser(dbWrapper)

	app.RunWithGracefulShutdown()
}
