package Presistence

import "Seller/Seller.Domain/Entities"

type IUserRepository interface {
	GetById(id int) Entities.User
	Create(entity any) Entities.User
	Update(entity any) bool
	Delete(entity any) bool
}
