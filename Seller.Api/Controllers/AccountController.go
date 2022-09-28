package Controller

import (
	"Seller/Seller.Application/Contract/Features/Account/Commands/Register"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	IUserRepository Presistence.IUserRepository
	SmsService      Infrastructure.ISMSService
}

// @Summary Register an User
// @Tags Register
// @Success 200 {object} Common.ResponseModel[string]
// @Param register body Register.RegisterCommand true "RegisterData"
// @Router /Register [post]
// @Security bearer
func (controller AccountController) RegisterUser(engine *gin.Engine) {
	engine.POST("/Register", func(c *gin.Context) {

		//jsonData, _ := ioutil.ReadAll(c.Request.Body)

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
}
