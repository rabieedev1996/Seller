package Config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnectionString = "host=192.168.179.130 user=postgres password=123456 dbname=Seller port=5432"

func GetDatabaseConnection() *gorm.DB {
	db, _ := gorm.Open(postgres.Open(DbConnectionString), &gorm.Config{})
	return db
}

type DatabaseConnection struct {
	Context *gorm.DB
}

func (r DatabaseConnection) Init() {
	r.Context = GetDatabaseConnection()
}
