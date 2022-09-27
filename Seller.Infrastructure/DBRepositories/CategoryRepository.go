package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type CategoryRepository struct {
	Generic Presistence.IGenericRepository[Entities.Category]
}

func (r CategoryRepository) Init() {
}

func (r CategoryRepository) GetById(id int) Entities.Category {
	return r.Generic.GetById(id)
}

func (r CategoryRepository) Create(model *Entities.Category) *Entities.Category {
	r.Generic.Create(model)
	return model
}
func (r CategoryRepository) Update(model *Entities.Category, id int) bool {
	return r.Generic.Update(model, id)
}

func (r CategoryRepository) Delete(model *Entities.Category) bool {
	return r.Generic.Delete(model)
}
