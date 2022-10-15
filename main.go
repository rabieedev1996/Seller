package main

import (
	GinHandler "Seller/Seller.Api/GinHandlers"
	_ "Seller/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Seller Api Documentation
// @version 1.0.0// @host localhost:8080
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
func main() {

	r := gin.Default()

	accountGroup := r.Group("/Account")
	categoriesGroup := r.Group("/Categories")
	productsGroup := r.Group("/Products")

	GinHandler.AccountHandler{
		Engine: accountGroup,
	}.RegisterUser().ForgetPassword().ActivateUser()

	GinHandler.CategoryHanlder{
		Engine: categoriesGroup,
	}.GetCategories().GetCategoryProducts()

	GinHandler.ProductHandler{
		Engine: productsGroup,
	}.GetProductDetail()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
