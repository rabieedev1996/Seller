package Controller

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Account/Category/Queries/GetCategories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	ICategoryRepository Presistence.ICategoryRepository
	Engine              *gin.RouterGroup
}

// @Summary Register an User
// @Tags Account
// @Success 200 {object} Common.ResponseModel[GetCategories.GetCategoriesVm]
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
		c.BindJSON(&requestQuery)
		handler := GetCategories.GetCategoriesQueryHandler{
			ICategoryRepository: controller.ICategoryRepository,
		}
		bussinesResult := handler.CommandHandler(requestQuery)
		c.JSON(http.StatusOK, bussinesResult)
	})
	return &controller
}
