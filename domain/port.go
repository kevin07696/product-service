package domain

type InventoryServicer interface {
	GetInventory(productID uint) (uint, StatusCode)
}
