package DBRepositories

import (
	"Seller/Config/InjectionConfig"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"reflect"
)

type UserRepository struct {
	Generic Presistence.IGenericRepository[Entities.User]
}

func (r UserRepository) Init(guid string) {
	r.Generic = InjectionConfig.PrepareObject[Presistence.IGenericRepository[Entities.User]](reflect.TypeOf(r.Generic), guid)
}

func (r UserRepository) GetById(id int) Entities.User {
	return r.Generic.GetById(id)
}

func (r UserRepository) Create(model Entities.User) Entities.User {
	r.Generic.Create(model)
	return model
}
func (r UserRepository) Update(model Entities.User, id int) bool {
	return r.Generic.Update(model)
}

func (r UserRepository) Delete(model Entities.User) bool {
	return r.Generic.Delete(model)
}
