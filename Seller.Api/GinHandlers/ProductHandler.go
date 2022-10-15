package GinHandler

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Engine *gin.RouterGroup
}

func (handler ProductHandler) GetController(c *gin.Context) *Controller.ProductController {
	db := Config.GetDatabaseConnection()
	controller := Controller.ProductController{
		ICategoryRepository: DBRepositories.CategoryRepository{
			Generic: Common.GetGenericRepository[Entities.Category](db),
		},
		ISelectableValueRepository: DBRepositories.SelectableValueRepository{
			Generic: Common.GetGenericRepository[Entities.SelectableValue](db),
		},
		IWarehouseRepository: DBRepositories.WarehouseRepository{
			Generic: Common.GetGenericRepository[Entities.Warehouse](db),
		},
		IProductRepository: DBRepositories.ProductRepository{
			Generic: Common.GetGenericRepository[Entities.Product](db),
		},
		Context: c,
	}

	return &controller
}

// @Summary Get Product Detail
// @Tags Products
// @Success 200 {object} Common.ResponseModel[GetProductDetail.GetProductQueryVM]
// @Router /Products/GetProductDetail [Get]
// @Param   productId   query      int  true  "productId"
// @Security bearer
func (handler ProductHandler) GetProductDetail() *ProductHandler {
	handler.Engine.GET("/GetProductDetail", func(c *gin.Context) {
		handler.GetController(c).GetProductDetail()
	})
	return &handler
}
