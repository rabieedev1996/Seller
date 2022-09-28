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

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Controller.AccountController{
		IUserRepository: iUserRepository,
		SmsService:      iSMSService,
	}.RegisterUser(r)
	r.Run()
}
