package Presistence

import (
	"Seller/Seller.Domain/Entities"
)

type ISelectableValueRepository interface {
	GetById(id int) Entities.SelectableValue
	GetAll() []Entities.SelectableValue
	Create(entity *Entities.SelectableValue) *Entities.SelectableValue
	Update(entity *Entities.SelectableValue, id any) bool
	Delete(entity *Entities.SelectableValue) bool
	GetProductSelectables(productId int) []Entities.SelectableValue
}
