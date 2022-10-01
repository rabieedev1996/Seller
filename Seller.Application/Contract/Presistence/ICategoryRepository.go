package Presistence

import (
	"Seller/Seller.Domain/Entities"
)

type ICategoryRepository interface {
	GetById(id string) Entities.Category
	GetAll() []Entities.Category
	Create(entity *Entities.Category) *Entities.Category
	Update(entity *Entities.Category, id any) bool
	Delete(entity *Entities.Category) bool
	GetCategoryChildsRecursive(id int) []Entities.Category
	GetCategoryProducts(categoryId int, start int, count int, sortType string) []Entities.Product
}
