package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type WarehouseRepository struct {
	Generic Presistence.IGenericRepository[Entities.Warehouse]
}

func (r WarehouseRepository) Init(guid string) {
}
func (r WarehouseRepository) GetAll() []Entities.Warehouse {
	return r.Generic.GetAll()
}
func (r WarehouseRepository) GetById(id int) Entities.Warehouse {
	return r.Generic.GetById(id)
}

func (r WarehouseRepository) Create(model *Entities.Warehouse) *Entities.Warehouse {
	r.Generic.Create(model)
	return model
}
func (r WarehouseRepository) Update(model *Entities.Warehouse, id any) bool {
	return r.Generic.Update(model, id)
}

func (r WarehouseRepository) Delete(model *Entities.Warehouse) bool {
	return r.Generic.Delete(model)
}

func (r WarehouseRepository) GetByProductId(id int) []Entities.Warehouse {
	result := []Entities.Warehouse{}
	r.Generic.GetContext().Table("Warehouse").Where("product_id=?", id).Find(&result)
	return result
}
