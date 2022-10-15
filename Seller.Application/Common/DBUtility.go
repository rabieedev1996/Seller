package Common

import (
	"Seller/Seller.Infrastructure/DBRepositories"
	"gorm.io/gorm"
)

func GetGenericRepository[entity any](db *gorm.DB) DBRepositories.GenericRepository[entity] {
	var genericRepository DBRepositories.GenericRepository[entity]
	genericRepository = DBRepositories.GenericRepository[entity]{
		Context: db,
	}
	return genericRepository
}
