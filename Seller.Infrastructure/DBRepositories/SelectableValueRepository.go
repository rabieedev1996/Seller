package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type SelectableValueRepository struct {
	Generic Presistence.IGenericRepository[Entities.SelectableValue]
}

func (r SelectableValueRepository) Init(guid string) {
}

func (r SelectableValueRepository) GetById(id int) Entities.SelectableValue {
	return r.Generic.GetById(id)
}

func (r SelectableValueRepository) Create(model Entities.SelectableValue) Entities.SelectableValue {
	r.Generic.Create(model)
	return model
}
func (r SelectableValueRepository) Update(model Entities.SelectableValue, id int) bool {
	return r.Generic.Update(model, id)
}

func (r SelectableValueRepository) Delete(model Entities.SelectableValue) bool {
	return r.Generic.Delete(model)
}
