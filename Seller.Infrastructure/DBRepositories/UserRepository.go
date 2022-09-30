package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type UserRepository struct {
	Generic Presistence.IGenericRepository[Entities.User]
}

func (r UserRepository) Init(guid string) {
	//r.Generic = *InjectionConfig.PrepareObject[Presistence.IGenericRepository[Entities.User]](reflect.TypeOf(r.Generic), guid)
}

func (r UserRepository) GetById(id string) Entities.User {
	return r.Generic.GetById(id)
}

func (r UserRepository) Create(model *Entities.User) *Entities.User {
	r.Generic.Create(model)
	return model
}
func (r UserRepository) Update(model *Entities.User, id any) bool {
	return r.Generic.Update(model, id)
}

func (r UserRepository) Delete(model *Entities.User) bool {
	return r.Generic.Delete(model)
}

func (r UserRepository) GetByUsername(username string) *Entities.User {
	var result Entities.User
	r.Generic.GetContext().Raw("select * from GetUserByUsername('" + username + "')").Scan(&result)
	return &result
}
