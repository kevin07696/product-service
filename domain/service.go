package domain

import "github.com/kevin07696/produce-service/generated"

type ProductService struct {
	repo      ProductRepositor
}

func NewProductService(repo ProductRepositor) ProductService {
	return ProductService{
		repo:      repo,
	}
}

func (s ProductService) GetProductDetail(request *generated.ProductDetailRequest) (productDetail ProductDetail, status StatusCode) {
	id := uint(request.ID)

	productDetail, status = s.repo.ReadProductDetail(id)
	if status > 0 {
		return
	}

	return
}

