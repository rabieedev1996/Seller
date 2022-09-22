package Presistence

type IGenericRepository[entity any] interface {
	GetById(id int) entity
	Create(entity any) entity
	Update(entity any) bool
	Delete(entity any) bool
}
