package GetProductDetail

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"github.com/devfeel/mapper"
)

type GetProductQueryHandler struct {
	CategoryRepository        Presistence.ICategoryRepository
	ProductRepository         Presistence.IProductRepository
	SelectableValueRepository Presistence.ISelectableValueRepository
	WarehouseRepository       Presistence.IWarehouseRepository
}

func (handler GetProductQueryHandler) HandlerFunc(query GetProductQuery) Common.ResponseModel[GetProductQueryVM] {

	product := handler.ProductRepository.GetById(query.ProductId)
	productCategory := handler.CategoryRepository.GetById(product.Category_Id)

	//productVm := GetProductQueryVM{
	//	Id:         product.Id,
	//	CategoryId: product.Category_Id,
	//	Title:      product.Title,
	//	Summary:    product.Summary,
	//	MinPrice:   product.Min_Price,
	//	MaxPrice:   product.Max_Price,
	//	LikeCount:  product.Like_Count,
	//	RankNumber: product.RankNumber,
	//}
	productVm := GetProductQueryVM{}
	mapper.AutoMapper(&product, &productVm)
	productVm = handler.fillProperties(query, &productVm)
	productVm = handler.fillSelectables(query, productCategory, &productVm)

	return (Common.ResponseModel[GetProductQueryVM]{}).ApiResponse(productVm, Common.SUCCESS)
}

func (handler GetProductQueryHandler) fillProperties(query GetProductQuery, productVm *GetProductQueryVM) GetProductQueryVM {
	productProperties := handler.ProductRepository.GetProductProperties(query.ProductId)
	for _, property := range productProperties {
		propertyVm := ProductPropertiesVM{}
		mapper.Mapper(&property, &propertyVm)
		productVm.Properties = append(productVm.Properties, propertyVm)
	}
	return *productVm
}

func (handler GetProductQueryHandler) fillSelectables(query GetProductQuery, productCategory Entities.Category, productVm *GetProductQueryVM) GetProductQueryVM {
	productSelectables := handler.SelectableValueRepository.GetProductSelectables(productCategory.Id)
	warehouse := handler.WarehouseRepository.GetByProductId(query.ProductId)
	productVm.Selectables = ProductSelectablesVM{
		Title: productCategory.Title,
	}
	existCount := 0
	for _, item := range productSelectables {
		for _, warehouseItem := range warehouse {
			if item.Id == warehouseItem.Selectable_Prop_Id {
				productVm.Selectables.Values = append(productVm.Selectables.Values, ProductSelectablesValuesVM{
					ValueTitle: item.Prop_Value_Title,
					ValueKey:   item.Prop_Value,
					ExistCount: warehouseItem.Exist_Count,
				})
				existCount += warehouseItem.Exist_Count
			}
		}
	}
	productVm.ExistCount = existCount

	return *productVm
}
