package main

import (
	GinHandler "Seller/Seller.Api/GinHandlers"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Features/Common/Command/UploadFileCommand"
	_ "Seller/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Seller Api Documentation
// @version 1.0.0// @host localhost:8080
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
func main() {

	r := gin.Default()
	r.Use(globalRecover)
	accountGroup := r.Group("/Account")
	categoriesGroup := r.Group("/Categories")
	productsGroup := r.Group("/Products")
	commonsGroup := r.Group("/Common")

	GinHandler.AccountHandler{
		Engine: accountGroup,
	}.RegisterUser().ForgetPassword().ActivateUser()

	GinHandler.CategoryHanlder{
		Engine: categoriesGroup,
	}.GetCategories().GetCategoryProducts()

	GinHandler.ProductHandler{
		Engine: productsGroup,
	}.GetProductDetail()

	GinHandler.CommonHandler{
		Engine: commonsGroup,
	}.UploadFile()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			result := (Common.ResponseModel[UploadFileCommand.UploadFileVm]{}).ApiResponse(nil, Common.STATUS_Exception)
			c.JSON(http.StatusOK, result)
		}
	}(c)
	c.Next()
}
