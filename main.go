package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func main() {

	loadProducts("products.json")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	productsGroup := r.Group("/products")
	{
		//productsGroup.GET("/:id", getByID)            // path param
		productsGroup.GET("/search", Search) // query param: products/search?priceGt=100&isPublished=true
	}

	r.Run(":8080")
}

func Search(c *gin.Context) {

	query := c.Query("priceGt")
	priceGt, err := strconv.ParseFloat(query, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Price",
		})
		return
	}

	list := []Product{}
	for _, product := range products {
		if product.Price > priceGt {
			list = append(list, product)
		}
	}

	c.JSON(http.StatusOK, list)
}

func loadProducts(path string) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &products); err != nil {
		panic(err)
	}
}
