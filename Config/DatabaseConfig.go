package Config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnectionString = "host=localhost user=postgres password=123456 dbname=MyDB port=5432"

func GetDatabaseConnection() *gorm.DB {
	db, _ := gorm.Open(postgres.Open(DbConnectionString), &gorm.Config{})
	return db
}
