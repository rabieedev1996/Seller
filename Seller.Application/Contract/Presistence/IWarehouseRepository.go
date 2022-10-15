package Presistence

import (
	"Seller/Seller.Domain/Entities"
)

type IWarehouseRepository interface {
	GetById(id int) Entities.Warehouse
	GetAll() []Entities.Warehouse
	Create(entity *Entities.Warehouse) *Entities.Warehouse
	Update(entity *Entities.Warehouse, id any) bool
	Delete(entity *Entities.Warehouse) bool
	GetByProductId(id int) []Entities.Warehouse
}
