package main

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"Seller/Seller.Infrastructure/Service"
	"github.com/gin-gonic/gin"
)

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
	Controller.AccountController{
		IUserRepository: iUserRepository,
		SmsService:      iSMSService,
	}.RegisterUser(r)
	r.Run()
}
