package GetCategories

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"github.com/devfeel/mapper"
)

type GetCategoriesQueryHandler struct {
	ICategoryRepository Presistence.ICategoryRepository
}

func (handler GetCategoriesQueryHandler) HandlerFunc(query GetCategoriesQuery) Common.ResponseModel[[]GetCategoriesVm] {
	var categories []Entities.Category
	if query.ParentId != 0 {
		categories = handler.ICategoryRepository.GetCategoryChildsRecursive(query.ParentId)
	} else {
		categories = handler.ICategoryRepository.GetAll()
	}

	resultData := []GetCategoriesVm{}
	for _, item := range categories {
		categoryVm := GetCategoriesVm{}
		mapper.AutoMapper(&item, &categoryVm)
		resultData = append(resultData, categoryVm)
	}
	return (Common.ResponseModel[[]GetCategoriesVm]{}).ApiResponse(&resultData, Common.STATUS_SUCCESS)
}
