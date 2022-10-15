package Controller

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Product/Queries/GetProductDetail"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	ICategoryRepository        Presistence.ICategoryRepository
	IProductRepository         Presistence.IProductRepository
	ISelectableValueRepository Presistence.ISelectableValueRepository
	IWarehouseRepository       Presistence.IWarehouseRepository
	Context                    *gin.Context
}

func (controller ProductController) GetProductDetail() *ProductController {
	productIdStr, _ := controller.Context.GetQuery("productId")
	productId, _ := strconv.Atoi(productIdStr)

	var requestQuery = GetProductDetail.GetProductQuery{
		ProductId: productId,
	}
	//json.Unmarshal(jsonData, requestBody)
	handler := GetProductDetail.GetProductQueryHandler{
		CategoryRepository:        controller.ICategoryRepository,
		ProductRepository:         controller.IProductRepository,
		SelectableValueRepository: controller.ISelectableValueRepository,
		WarehouseRepository:       controller.IWarehouseRepository,
	}
	bussinesResult := handler.HandlerFunc(requestQuery)

	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}
