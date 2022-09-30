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

func (handler GetCategoriesQueryHandler) CommandHandler(command GetCategoriesQuery) Common.ResponseModel[[]GetCategoriesVm] {
	var categories []Entities.Category
	if command.ParentId != 0 {
		categories = handler.ICategoryRepository.GetCategoryChildsRecursive(command.ParentId)
	} else {
		categories = handler.ICategoryRepository.GetAll()
	}

	resultData := []GetCategoriesVm{}
	for _, item := range categories {
		categoryVm := GetCategoriesVm{}
		mapper.AutoMapper(&item, &categoryVm)
		resultData = append(resultData, categoryVm)
	}
	return (Common.ResponseModel[[]GetCategoriesVm]{}).ApiResponse(resultData, Common.SUCCESS)
}
