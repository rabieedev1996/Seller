package Controller

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Category/Queries/GetCategories"
	"Seller/Seller.Application/Features/Category/Queries/GetCategoryProducts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	ICategoryRepository Presistence.ICategoryRepository
	Context             *gin.Context
}

func (controller CategoryController) GetCategories() *CategoryController {
	parentIdString, _ := controller.Context.GetQuery("parentId")
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
	bussinesResult := handler.HandlerFunc(requestQuery)
	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}

func (controller CategoryController) GetCategoryProducts() *CategoryController {
	var requestBody = GetCategoryProducts.GetCategoryProductsQuery{}
	//json.Unmarshal(jsonData, requestBody)
	controller.Context.BindJSON(&requestBody)

	//json.Unmarshal(jsonData, requestBody)
	handler := GetCategoryProducts.GetCategoryProductsHandler{
		ICategoryRepository: controller.ICategoryRepository,
	}
	bussinesResult := handler.CommandHandler(requestBody)
	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}
