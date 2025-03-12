package domain

import "gorm.io/datatypes"

type ProductWriter interface {
	WriteProduct(product Product) StatusCode
	UpdateProduct(product Product) StatusCode
	DeleteProduct(productID datatypes.UUID) StatusCode
}
