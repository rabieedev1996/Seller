package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type CategoryPropertyRepository struct {
	Generic Presistence.IGenericRepository[Entities.CategoryProperty]
}

func (r CategoryPropertyRepository) Init() {
}

func (r CategoryPropertyRepository) GetById(id int) Entities.CategoryProperty {
	return r.Generic.GetById(id)
}

func (r CategoryPropertyRepository) Create(model *Entities.CategoryProperty) *Entities.CategoryProperty {
	r.Generic.Create(model)
	return model
}
func (r CategoryPropertyRepository) Update(model *Entities.CategoryProperty, id int) bool {
	return r.Generic.Update(model, id)
}

func (r CategoryPropertyRepository) Delete(model *Entities.CategoryProperty) bool {
	return r.Generic.Delete(model)
}
