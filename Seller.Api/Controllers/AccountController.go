package Controller

import (
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
	Context         *gin.Context
}

func (controller AccountController) RegisterUser() *AccountController {
	var requestBody = Register.RegisterCommand{}
	//json.Unmarshal(jsonData, requestBody)
	controller.Context.BindJSON(&requestBody)
	handler := Register.RegisterCommandHandler{
		IUserRepository: controller.IUserRepository,
		SmsService:      controller.SmsService,
	}
	bussinesResult := handler.HandlerFunc(requestBody)
	controller.Context.JSON(http.StatusOK, bussinesResult)
	return &controller
}

func (controller AccountController) ActivateUser() *AccountController {

	userId := Common.GetClaimValuesFromGin(controller.Context).UserId
	var requestBody = ActivationByCode.ActivationByCodeCommand{}
	//json.Unmarshal(jsonData, requestBody)
	controller.Context.BindJSON(&requestBody)
	handler := ActivationByCode.ActivationByCodeCommandHandler{
		IUserRepository: controller.IUserRepository,
		SmsService:      controller.SmsService,
	}
	bussinesResult := handler.HandlerFunc(requestBody, userId)
	controller.Context.JSON(http.StatusOK, bussinesResult)

	return &controller
}

func (controller AccountController) ForgetPassword() *AccountController {

	//jsonData, _ := ioutil.ReadAll(c.Request.Body)

	var requestBody = ForgetPassword.ForgetPasswordCommand{}
	//json.Unmarshal(jsonData, requestBody)
	controller.Context.BindJSON(&requestBody)
	handler := ForgetPassword.ForgetPasswordCommandHandler{
		IUserRepository: controller.IUserRepository,
		SmsService:      controller.SmsService,
	}
	bussinesResult := handler.HandlerFunc(requestBody)
	controller.Context.JSON(http.StatusOK, bussinesResult)

	return &controller
}
