package main

import (
	"database/sql"

	"github.com/AlejaMarin/go-api/cmd/server/handler"
	"github.com/AlejaMarin/go-api/internal/product"
	"github.com/AlejaMarin/go-api/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// storage := store.NewJsonStore("../../products.json")

	db, err := sql.Open("mysql", "user1:secret_password@/my_db")
	if err != nil {
		panic(err)
	}
	storage := store.NewSqlStore(db)

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
