package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type PropertiesValueRepository struct {
	Generic Presistence.IGenericRepository[Entities.PropertiesValue]
}

func (r PropertiesValueRepository) Init(guid string) {
}

func (r PropertiesValueRepository) GetById(id int) Entities.PropertiesValue {
	return r.Generic.GetById(id)
}

func (r PropertiesValueRepository) Create(model *Entities.PropertiesValue) *Entities.PropertiesValue {
	r.Generic.Create(model)
	return model
}
func (r PropertiesValueRepository) Update(model *Entities.PropertiesValue, id int) bool {
	return r.Generic.Update(model, id)
}

func (r PropertiesValueRepository) Delete(model *Entities.PropertiesValue) bool {
	return r.Generic.Delete(model)
}
