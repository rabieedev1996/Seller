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

func (handler GetCategoryProductsHandler) CommandHandler(query GetCategoryProductsQuery) Common.ResponseModel[[]GetCategoryProductsVm] {
	if query.Count == 0 {
		query.Count = 10
	}
	if query.SortType == "" {
		query.SortType = "dateDesc"
	}

	var categories []Entities.Product
	categories = handler.ICategoryRepository.GetCategoryProducts(query.CategoryId, query.Start, query.Count, query.SortType)

	resultData := []GetCategoryProductsVm{}
	for _, item := range categories {
		categoryVm := GetCategoryProductsVm{}
		mapper.AutoMapper(&item, &categoryVm)
		resultData = append(resultData, categoryVm)
	}
	return (Common.ResponseModel[[]GetCategoryProductsVm]{}).ApiResponse(resultData, Common.SUCCESS)
}
