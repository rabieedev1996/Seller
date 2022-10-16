package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
)

type FileRepository struct {
	Generic Presistence.IGenericRepository[Entities.File]
}

func (r FileRepository) Init() {
}
func (r FileRepository) GetAll() []Entities.File {
	return r.Generic.GetAll()
}
func (r FileRepository) GetById(id int) Entities.File {
	return r.Generic.GetById(id)
}

func (r FileRepository) Create(model *Entities.File) *Entities.File {
	r.Generic.Create(model)
	return model
}
func (r FileRepository) Update(model *Entities.File, id int) bool {
	return r.Generic.Update(model, id)
}

func (r FileRepository) Delete(model *Entities.File) bool {
	return r.Generic.Delete(model)
}
