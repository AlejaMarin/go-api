package domain

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type ProductPatchRequest struct {
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
