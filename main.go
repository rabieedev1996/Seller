package main

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"Seller/Seller.Infrastructure/Service"
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

	db := Config.GetDatabaseConnection()

	var iSMSService Infrastructure.ISMSService
	iSMSService = Service.FarazSMSService{}

	var userGenericRepository DBRepositories.GenericRepository[Entities.User]
	userGenericRepository = DBRepositories.GenericRepository[Entities.User]{
		Context: db,
	}
	var iUserRepository Presistence.IUserRepository
	iUserRepository = DBRepositories.UserRepository{
		Generic: userGenericRepository,
	}

	var categoryGenericRepository DBRepositories.GenericRepository[Entities.Category]
	categoryGenericRepository = DBRepositories.GenericRepository[Entities.Category]{
		Context: db,
	}
	var iCategoryRepository Presistence.ICategoryRepository
	iCategoryRepository = DBRepositories.CategoryRepository{
		Generic: categoryGenericRepository,
	}

	r := gin.Default()

	accountGroup := r.Group("/Account")
	categoriesGroup := r.Group("/Categories")
	Controller.AccountController{
		IUserRepository: iUserRepository,
		SmsService:      iSMSService,
		Engine:          accountGroup,
	}.RegisterUser().ForgetPassword().ActivateUser()

	Controller.CategoryController{
		ICategoryRepository: iCategoryRepository,
		Engine:              categoriesGroup,
	}.GetCategories().GetCategoryProducts()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
