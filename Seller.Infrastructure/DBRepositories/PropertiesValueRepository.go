package DBRepositories

import (
	"Seller/Config/InjectionConfig"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"reflect"
)

type PropertiesValueRepository struct {
	Generic Presistence.IGenericRepository[Entities.PropertiesValue]
}

func (r PropertiesValueRepository) Init(guid string) {
	r.Generic = InjectionConfig.PrepareObject[Presistence.IGenericRepository[Entities.PropertiesValue]](reflect.TypeOf(r.Generic), guid)
}

func (r PropertiesValueRepository) GetById(id int) Entities.PropertiesValue {
	return r.Generic.GetById(id)
}

func (r PropertiesValueRepository) Create(model Entities.PropertiesValue) Entities.PropertiesValue {
	r.Generic.Create(model)
	return model
}
func (r PropertiesValueRepository) Update(model Entities.PropertiesValue, id int) bool {
	return r.Generic.Update(model)
}

func (r PropertiesValueRepository) Delete(model Entities.PropertiesValue) bool {
	return r.Generic.Delete(model)
}
