package GinHandler

import (
	"Seller/Config"
	Controller "Seller/Seller.Api/Controllers"
	"Seller/Seller.Api/Middlewares"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"Seller/Seller.Infrastructure/Service"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	IUserRepository Presistence.IUserRepository
	SmsService      Infrastructure.ISMSService
	Engine          *gin.RouterGroup
}

// @Summary Register an User
// @Tags Account
// @Success 200 {object} Common.ResponseModel[string]
// @Param register body Register.RegisterCommand true "RegisterData"
// @Router /Account/Register [post]
// @Security bearer
func (hanlder AccountHandler) RegisterUser() *AccountHandler {
	hanlder.Engine.POST("/Register", func(c *gin.Context) {
		hanlder.createController(c).RegisterUser()
	})
	return &hanlder
}

// @Summary Activate User by Code
// @Tags Account
// @Success 200 {object} Common.ResponseModel[string]
// @Param ActivateUser body ActivationByCode.ActivationByCodeCommand true "ActivationData"
// @Router /Account/ActivateUser [post]
// @Security bearer
func (hanlder AccountHandler) ActivateUser() *AccountHandler {
	hanlder.Engine.Use(Middlewares.Authorize())
	{
		hanlder.Engine.POST("/ActivateUser", func(c *gin.Context) {
			hanlder.createController(c).RegisterUser()
		})
	}
	return &hanlder
}

// @Summary Register an User
// @Tags Account
// @Success 200 {object} Common.ResponseModel[string]
// @Param ForgetPassword body ForgetPassword.ForgetPasswordCommand true "ForgetPassword"
// @Router /Account/ForgetPassword [post]
// @Security bearer
func (hanlder AccountHandler) ForgetPassword() *AccountHandler {
	hanlder.Engine.POST("/ForgetPassword", func(c *gin.Context) {
		hanlder.createController(c).ForgetPassword()
	})
	return &hanlder
}

func (hanlder AccountHandler) createController(c *gin.Context) Controller.AccountController {
	var iSMSService Infrastructure.ISMSService
	iSMSService = Service.FarazSMSService{}
	db := Config.GetDatabaseConnection()

	controller := Controller.AccountController{
		Context: c,
		IUserRepository: DBRepositories.UserRepository{
			Generic: Common.GetGenericRepository[Entities.User](db),
		},
		SmsService: iSMSService,
	}
	return controller
}
