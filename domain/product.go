package domain

import "time"

type Product struct {
	ID                 string    `json:"id"`
	ProductName        string    `json:"product_name"`
	ProductMaterial    string    `json:"product_material"`
	ProductDescription string    `json:"product_description"`
	ProductPrice       string    `json:"product_price"`
	ProductImages      string    `json:"product_images"`
	CreatedAt          time.Time `json:"created_at"`
}

type ProductRepository interface {
	CheckProductByID(id int) error
}