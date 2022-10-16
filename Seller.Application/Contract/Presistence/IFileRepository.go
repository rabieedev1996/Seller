package Presistence

import "Seller/Seller.Domain/Entities"

type IFileRepository interface {
	GetById(id int) Entities.File
	GetAll() []Entities.File
	Create(entity *Entities.File) *Entities.File
	Update(entity *Entities.File, id int) bool
	Delete(entity *Entities.File) bool
}
