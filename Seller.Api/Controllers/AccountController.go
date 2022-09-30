package Controller

import (
	"Seller/Seller.Api/Middlewares"
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Application/Features/Account/Commands/ActivationByCode"
	"Seller/Seller.Application/Features/Account/Commands/ForgetPassword"
	"Seller/Seller.Application/Features/Account/Commands/Register"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
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
func (controller AccountController) RegisterUser() *AccountController {
	controller.Engine.POST("/Register", func(c *gin.Context) {
		var requestBody = Register.RegisterCommand{}
		//json.Unmarshal(jsonData, requestBody)
		c.BindJSON(&requestBody)
		handler := Register.RegisterCommandHandler{
			IUserRepository: controller.IUserRepository,
			SmsService:      controller.SmsService,
		}
		bussinesResult := handler.CommandHandler(requestBody)
		c.JSON(http.StatusOK, bussinesResult)
	})
	return &controller
}

// @Summary Activate User by Code
// @Tags Account
// @Success 200 {object} Common.ResponseModel[string]
// @Param ActivateUser body ActivationByCode.ActivationByCodeCommand true "ActivationData"
// @Router /Account/ActivateUser [post]
// @Security bearer
func (controller AccountController) ActivateUser() *AccountController {
	controller.Engine.Use(Middlewares.Authorize())
	{
		controller.Engine.POST("/ActivateUser", func(c *gin.Context) {
			userId := Common.GetClaimValuesFromGin(c).UserId
			var requestBody = ActivationByCode.ActivationByCodeCommand{}
			//json.Unmarshal(jsonData, requestBody)
			c.BindJSON(&requestBody)
			handler := ActivationByCode.ActivationByCodeCommandHandler{
				IUserRepository: controller.IUserRepository,
				SmsService:      controller.SmsService,
			}
			bussinesResult := handler.CommandHandler(requestBody, userId)
			c.JSON(http.StatusOK, bussinesResult)
		})
	}
	return &controller
}

// @Summary Register an User
// @Tags Account
// @Success 200 {object} Common.ResponseModel[string]
// @Param ForgetPassword body ForgetPassword.ForgetPasswordCommand true "ForgetPassword"
// @Router /Account/ForgetPassword [post]
// @Security bearer
func (controller AccountController) ForgetPassword() *AccountController {
	controller.Engine.POST("/ForgetPassword", func(c *gin.Context) {

		//jsonData, _ := ioutil.ReadAll(c.Request.Body)

		var requestBody = ForgetPassword.ForgetPasswordCommand{}
		//json.Unmarshal(jsonData, requestBody)
		c.BindJSON(&requestBody)
		handler := ForgetPassword.ForgetPasswordCommandHandler{
			IUserRepository: controller.IUserRepository,
			SmsService:      controller.SmsService,
		}
		bussinesResult := handler.CommandHandler(requestBody)
		c.JSON(http.StatusOK, bussinesResult)
	})
	return &controller
}
