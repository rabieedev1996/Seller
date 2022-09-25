package Presistence

type IGenericRepository[entity any] interface {
	GetById(id int) entity
	Create(model entity) entity
	Update(model entity, id int) bool
	Delete(model entity) bool
}
