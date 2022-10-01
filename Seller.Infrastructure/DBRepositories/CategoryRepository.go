package DBRepositories

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"strconv"
)

type CategoryRepository struct {
	Generic Presistence.IGenericRepository[Entities.Category]
}

func (r CategoryRepository) Init() {
}

func (r CategoryRepository) GetById(id string) Entities.Category {
	return r.Generic.GetById(id)
}
func (r CategoryRepository) GetAll() []Entities.Category {
	return r.Generic.GetAll()
}

func (r CategoryRepository) Create(model *Entities.Category) *Entities.Category {
	r.Generic.Create(model)
	return model
}
func (r CategoryRepository) Update(model *Entities.Category, id any) bool {
	return r.Generic.Update(model, id)
}

func (r CategoryRepository) Delete(model *Entities.Category) bool {
	return r.Generic.Delete(model)
}

func (r CategoryRepository) GetCategoryChildsRecursive(id int) []Entities.Category {
	var result []Entities.Category
	r.Generic.GetContext().Raw("select * from \"Category\" where \"id\" =any(public.getCategoryChildsRecursive(" + strconv.Itoa(id) + "))").Scan(&result)
	return result
}

func (r CategoryRepository) GetCategoryProducts(categoryId int, start int, count int, sortType string) []Entities.Product {
	var result []Entities.Product
	r.Generic.GetContext().Raw("select * from GetCategoryProducts(" + strconv.Itoa(categoryId) + "," + strconv.Itoa(start) + "," + strconv.Itoa(count) + ",'" + sortType + "')").Scan(&result)
	return result
}
