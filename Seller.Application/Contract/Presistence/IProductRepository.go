package Presistence

import (
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Domain/QueryModels"
)

type IProductRepository interface {
	GetById(id int) Entities.Product
	GetAll() []Entities.Product
	Create(entity *Entities.Product) *Entities.Product
	Update(entity *Entities.Product, id any) bool
	Delete(entity *Entities.Product) bool
	GetProductProperties(productId int) []QueryModels.GetProductPropertiesResult
}
