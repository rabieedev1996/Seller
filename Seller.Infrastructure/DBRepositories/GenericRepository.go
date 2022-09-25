package DBRepositories

import (
	"Seller/Config/InjectionConfig"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type GenericRepository[entity any] struct {
	Context *gorm.DB
}

func (r GenericRepository[entity]) Init(guid string) {
	databaseConnection := InjectionConfig.PrepareObject[DatabaseConnection](reflect.TypeOf(DatabaseConnection{}), guid)
	r.Context = databaseConnection.Context
}

func (r GenericRepository[entity]) GetById(id int) entity {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	var model entity
	r.Context.Table(titleArray[len(titleArray)-1]).First(model, id)
	return model
}

func (r GenericRepository[entity]) Create(model entity) entity {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Create(model)
	return model
}

func (r GenericRepository[entity]) Update(model entity, id int) bool {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Save(model)
	return true
}

func (r GenericRepository[entity]) Delete(model entity) bool {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Delete(model)
	return true
}
