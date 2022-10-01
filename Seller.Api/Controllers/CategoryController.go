package Controller

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Account/Category/Queries/GetCategories"
	"Seller/Seller.Application/Features/Account/Category/Queries/GetCategoryProducts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	ICategoryRepository Presistence.ICategoryRepository
	Engine              *gin.RouterGroup
}

// @Summary Get Categories List
// @Tags Categories
// @Success 200 {object} Common.ResponseModel[[]GetCategories.GetCategoriesVm]
// @Router /Categories/GetCategories [Get]
// @Param        parentId   query      int  false  "Parent Id"
// @Security bearer
func (controller CategoryController) GetCategories() *CategoryController {
	controller.Engine.GET("/GetCategories", func(c *gin.Context) {
		parentIdString, _ := c.GetQuery("parentId")
		var parentId int
		if parentIdString != "" {
			parentId, _ = strconv.Atoi(parentIdString)
		} else {
			parentId = 0
		}

		var requestQuery = GetCategories.GetCategoriesQuery{
			ParentId: parentId,
		}
		//json.Unmarshal(jsonData, requestBody)
		handler := GetCategories.GetCategoriesQueryHandler{
			ICategoryRepository: controller.ICategoryRepository,
		}
		bussinesResult := handler.CommandHandler(requestQuery)
		c.JSON(http.StatusOK, bussinesResult)
	})
	return &controller
}

// @Summary Get Category Products
// @Tags Categories
// @Success 200 {object} Common.ResponseModel[[]GetCategories.GetCategoriesVm]
// @Router /Categories/GetCategoryProducts [Post]
// @Param GetCategoryProduct body GetCategoryProducts.GetCategoryProductsQuery true "GetCategoryProduct"
// @Security bearer
func (controller CategoryController) GetCategoryProducts() *CategoryController {
	controller.Engine.POST("/GetCategoryProducts", func(c *gin.Context) {
		var requestBody = GetCategoryProducts.GetCategoryProductsQuery{}
		//json.Unmarshal(jsonData, requestBody)
		c.BindJSON(&requestBody)
		//json.Unmarshal(jsonData, requestBody)
		handler := GetCategoryProducts.GetCategoryProductsHandler{
			ICategoryRepository: controller.ICategoryRepository,
		}
		bussinesResult := handler.CommandHandler(requestBody)
		c.JSON(http.StatusOK, bussinesResult)
	})
	return &controller
}
