package ForgetPassword

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"strconv"
	"time"
)

type ForgetPasswordCommandHandler struct {
	IUserRepository Presistence.IUserRepository
	SmsService      Infrastructure.ISMSService
}

func (handler ForgetPasswordCommandHandler) HandlerFunc(command ForgetPasswordCommand) Common.ResponseModel[string] {
	user := handler.IUserRepository.GetByUsername(strconv.FormatInt(command.PhoneNumber, 10))
	user.Activation_Code = Common.GetRandomCode(100000, 999999)
	handler.IUserRepository.Update(user, user.Id)
	handler.SmsService.SendVerificationCode("0"+strconv.FormatInt(user.Phone_Number, 10), strconv.Itoa(user.Activation_Code))

	tokenExp := time.Now().Add(2 * time.Minute)
	token, _ := Common.GenerateToken(user.Id, tokenExp)

	return (Common.ResponseModel[string]{}).ApiResponse(token, Common.SUCCESS)
}
