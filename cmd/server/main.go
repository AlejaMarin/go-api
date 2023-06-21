package main

import (
	"github.com/AlejaMarin/go-api/cmd/server/handler"
	"github.com/AlejaMarin/go-api/internal/product"
	"github.com/AlejaMarin/go-api/pkg/middleware"
	"github.com/AlejaMarin/go-api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storage := store.NewStore("../../products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.GET("/consumer_price", productHandler.ConsumerPrice())
		products.POST("", middleware.TokenAuthMiddleware(), productHandler.Post())
		products.DELETE(":id", middleware.TokenAuthMiddleware(), productHandler.Delete())
		products.PATCH(":id", middleware.TokenAuthMiddleware(), productHandler.Patch())
		products.PUT(":id", middleware.TokenAuthMiddleware(), productHandler.Put())
	}
	r.Run(":8080")
}
