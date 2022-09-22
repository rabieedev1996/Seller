package DBRepositories

import (
	"Seller/Config/InjectionConfig"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"reflect"
)

type SelectableValueRepository struct {
	Generic Presistence.IGenericRepository[Entities.SelectableValue]
}

func (r SelectableValueRepository) Init() {
	r.Generic = InjectionConfig.PrepareObject[Presistence.IGenericRepository[Entities.SelectableValue]](reflect.TypeOf(r.Generic))
}

func (r SelectableValueRepository) GetById(id int) Entities.SelectableValue {
	return r.Generic.GetById(id)
}

func (r SelectableValueRepository) Create(model Entities.SelectableValue) Entities.SelectableValue {
	r.Generic.Create(model)
	return model
}
func (r SelectableValueRepository) Update(model Entities.SelectableValue, id int) bool {
	return r.Generic.Update(model)
}

func (r SelectableValueRepository) Delete(model Entities.SelectableValue) bool {
	return r.Generic.Delete(model)
}
