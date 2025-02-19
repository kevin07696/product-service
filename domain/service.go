package domain

import "github.com/kevin07696/produce-service/generated"

type ProductService struct {
	repo      ProductRepositor
	inventory InventoryServicer
}

func NewProductService(repo ProductRepositor, inventory InventoryServicer) ProductService {
	return ProductService{
		repo:      repo,
		inventory: inventory,
	}
}

func (s ProductService) GetProductDetail(request *generated.ProductDetailRequest) (productDetail ProductDetail, status StatusCode) {
	id := uint(request.ID)

	productDetail, status = s.repo.ReadProductDetail(id)
	if status > 0 {
		return
	}

	productDetail.Stock, status = s.inventory.GetInventory(id)
	return
}

