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
		productsGroup.GET("/search", Search) // query param: products/search?priceGt=100&isPublished=true
		productsGroup.GET("/productparams", AddProduct)
		productsGroup.GET("/:id", GetById)
		productsGroup.GET("/searchbyquantity", SearchByQuantity)
		productsGroup.GET("/buy", BuyProduct)
	}

	r.Run(":8080")
}

func AddProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id",
		})
		return
	}
	name := c.Query("name")
	quantity, err2 := strconv.Atoi(c.Query("quantity"))
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Quantity",
		})
		return
	}
	codeValue := c.Query("code_value")
	isP, err3 := strconv.ParseBool(c.Query("is_published"))
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Bool",
		})
		return
	}
	exp := c.Query("expiration")
	price, err4 := strconv.ParseFloat(c.Query("price"), 64)
	if err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Price",
		})
		return
	}

	p := Product{id, name, quantity, codeValue, isP, exp, price}
	products = append(products, p)

	c.JSON(http.StatusOK, p)
}

func GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id",
		})
		return
	}
	var encontrado Product
	for _, v := range products {
		if id == v.ID {
			encontrado = v
		}
	}
	if encontrado != (Product{}) {
		c.JSON(http.StatusOK, encontrado)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product Not Found",
		})
	}
}

func SearchByQuantity(c *gin.Context) {

	min, err := strconv.Atoi(c.Query("min"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Min",
		})
		return
	}
	max, err := strconv.Atoi(c.Query("max"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Max",
		})
		return
	}
	var lista []Product
	for _, v := range products {

		if v.Quantity > min && v.Quantity < max {
			lista = append(lista, v)
		}

	}
	if len(lista) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Products Not Found",
		})
	} else {
		c.JSON(http.StatusOK, lista)
	}

}

func BuyProduct(c *gin.Context) {

	code := c.Query("code_value")
	cant, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Quantity",
		})
		return
	}
	var encontrado Product
	for _, v := range products {
		if code == v.CodeValue {
			encontrado = v
		}
	}
	if encontrado != (Product{}) {
		c.JSON(http.StatusOK, gin.H{
			"nombre-producto": encontrado.Name,
			"cantidad":        cant,
			"precio-total":    float64(cant) * (encontrado.Price),
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product Not Found",
		})
	}

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
