package Presistence

import "Seller/Seller.Domain/Entities"

type IUserRepository interface {
	GetById(id int) Entities.User
	Create(entity *Entities.User) *Entities.User
	Update(entity *Entities.User, id any) bool
	Delete(entity *Entities.User) bool
	GetByUsername(userName string) *Entities.User
}
