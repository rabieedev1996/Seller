package ActivationByCode

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"time"
)

type ActivationByCodeCommandHandler struct {
	IUserRepository Presistence.IUserRepository
	SmsService      Infrastructure.ISMSService
}

func (handler ActivationByCodeCommandHandler) HandlerFunc(command ActivationByCodeCommand, userId string) Common.ResponseModel[string] {
	user := handler.IUserRepository.GetById(userId)
	if user.Activation_Code == command.ActivationCode {
		user.Is_Active = true
		handler.IUserRepository.Update(&user, user.Id)
	} else {
		return (Common.ResponseModel[string]{}).ApiResponse(nil, Common.STATUS_INVALID_ACTIVATION_CODE)
	}
	tokenExp := time.Now().AddDate(1, 0, 0)
	token, _ := Common.GenerateToken(user.Id, tokenExp)

	return (Common.ResponseModel[string]{}).ApiResponse(&token, Common.STATUS_SUCCESS)
}
