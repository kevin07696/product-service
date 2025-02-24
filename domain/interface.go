package domain

import "github.com/kevin07696/produce-service/generated"

type ProductRepositor interface {
	Migrate()
	ReadProductSummaries(request *generated.ProductSummaryRequest) ([]ProductSummary, StatusCode)
	ReadProductDetail(productID uint) (ProductDetail, StatusCode)
	WriteProducts() StatusCode
	ReadCategories() ([]string, StatusCode)
}
