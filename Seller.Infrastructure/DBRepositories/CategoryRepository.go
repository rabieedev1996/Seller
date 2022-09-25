package DBRepositories

import (
	"Seller/Config/InjectionConfig"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"reflect"
)

type CategoryRepository struct {
	Generic Presistence.IGenericRepository[Entities.Category]
}

func (r CategoryRepository) Init(guid string) {
	r.Generic = *InjectionConfig.PrepareObject[Presistence.IGenericRepository[Entities.Category]](reflect.TypeOf(r.Generic), guid)
}

func (r CategoryRepository) GetById(id int) Entities.Category {
	return r.Generic.GetById(id)
}

func (r CategoryRepository) Create(model Entities.Category) Entities.Category {
	r.Generic.Create(model)
	return model
}
func (r CategoryRepository) Update(model Entities.Category, id int) bool {
	return r.Generic.Update(model)
}

func (r CategoryRepository) Delete(model Entities.Category) bool {
	return r.Generic.Delete(model)
}
