package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type ProductRepository struct {
	Generic Presistence.IGenericRepository[Entities.Product]
}

func (r ProductRepository) Init(guid string) {
}

func (r ProductRepository) GetById(id int) Entities.Product {
	return r.Generic.GetById(id)
}

func (r ProductRepository) Create(model *Entities.Product) *Entities.Product {
	r.Generic.Create(model)
	return model
}
func (r ProductRepository) Update(model *Entities.Product, id int) bool {
	return r.Generic.Update(model, id)
}

func (r ProductRepository) Delete(model *Entities.Product) bool {
	return r.Generic.Delete(model)
}
