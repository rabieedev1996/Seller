package Presistence

import "gorm.io/gorm"

type IGenericRepository[entity any] interface {
	GetById(id int) entity
	Create(model *entity) *entity
	Update(model *entity, id any) bool
	Delete(model *entity) bool

	GetContext() *gorm.DB
}
