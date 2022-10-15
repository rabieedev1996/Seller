package GinHandler

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"github.com/gin-gonic/gin"
)

type CategoryHanlder struct {
	Engine *gin.RouterGroup
}

func (handler *CategoryHanlder) createController(c *gin.Context) Controller.CategoryController {
	db := Config.GetDatabaseConnection()
	categoryRepository := DBRepositories.CategoryRepository{
		Generic: Common.GetGenericRepository[Entities.Category](db),
	}

	controller := Controller.CategoryController{
		Context:             c,
		ICategoryRepository: categoryRepository,
	}
	return controller
}

// @Summary Get Categories List
// @Tags Categories
// @Success 200 {object} Common.ResponseModel[[]GetCategories.GetCategoriesVm]
// @Router /Categories/GetCategories [Get]
// @Param        parentId   query      int  false  "Parent Id"
// @Security bearer
func (handler CategoryHanlder) GetCategories() *CategoryHanlder {
	handler.Engine.GET("/GetCategories", func(c *gin.Context) {
		handler.createController(c).GetCategories()
	})
	return &handler
}

// @Summary Get Category Products
// @Tags Categories
// @Success 200 {object} Common.ResponseModel[[]GetCategories.GetCategoriesVm]
// @Router /Categories/GetCategoryProducts [Post]
// @Param GetCategoryProduct body GetCategoryProducts.GetCategoryProductsQuery true "GetCategoryProduct"
// @Security bearer
func (handler CategoryHanlder) GetCategoryProducts() *CategoryHanlder {
	handler.Engine.POST("/GetCategoryProducts", func(c *gin.Context) {
		handler.createController(c).GetCategoryProducts()
	})
	return &handler
}
