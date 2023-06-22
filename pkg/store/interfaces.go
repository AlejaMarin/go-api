package store

import "github.com/AlejaMarin/go-api/internal/domain"

type StoreInterface interface {
	// Devuelve lista de todos los productos
	GetAll() ([]domain.Product, error)
	// Read devuelve un producto por su id
	Read(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(product domain.Product) error
	// Update actualiza un producto
	Update(product domain.Product) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
