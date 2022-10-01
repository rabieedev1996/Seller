package GetCategoryProducts

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"github.com/devfeel/mapper"
)

type GetCategoryProductsHandler struct {
	ICategoryRepository Presistence.ICategoryRepository
}

func (handler GetCategoryProductsHandler) CommandHandler(command GetCategoryProductsQuery) Common.ResponseModel[[]GetCategoryProductsVm] {
	var categories []Entities.Product
	categories = handler.ICategoryRepository.GetCategoryProducts(command.CategoryId, command.Start, command.Count, command.SortType)

	resultData := []GetCategoryProductsVm{}
	for _, item := range categories {
		categoryVm := GetCategoryProductsVm{}
		mapper.AutoMapper(&item, &categoryVm)
		resultData = append(resultData, categoryVm)
	}
	return (Common.ResponseModel[[]GetCategoryProductsVm]{}).ApiResponse(resultData, Common.SUCCESS)
}
