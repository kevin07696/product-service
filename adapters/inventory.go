package adapters

import "github.com/kevin07696/produce-service/domain"

type InventoryService struct {}

func (s InventoryService) GetInventory(productID uint) (uint, domain.StatusCode) {
	return 1, domain.StatusOK
}