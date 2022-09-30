package Register

import (
	"Seller/Seller.Application/Common"
	"Seller/Seller.Application/Contract/Infrastructure"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"github.com/google/uuid"
	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type RegisterCommandHandler struct {
	IUserRepository Presistence.IUserRepository
	SmsService      Infrastructure.ISMSService
}

func (handler RegisterCommandHandler) CommandHandler(command RegisterCommand) Common.ResponseModel[string] {
	user := handler.IUserRepository.GetByUsername(strconv.FormatInt(command.PhoneNumber, 10))
	hash, _ := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)

	jalaliDate := ptime.Now()
	numericJalaliDate, _ := strconv.Atoi(jalaliDate.Format("yyyyMMdd"))
	if user.Id == "" {
		user = &Entities.User{
			Id:                uuid.New().String(),
			Phone_Number:      command.PhoneNumber,
			Username:          strconv.FormatInt(command.PhoneNumber, 10),
			Is_Active:         false,
			Password_Hash:     string(hash),
			Activation_Code:   Common.GetRandomCode(100000, 999999),
			Created_At:        time.Now(),
			Jalali_Created_At: numericJalaliDate,
		}
		handler.IUserRepository.Create(user)
	} else {
		user.Is_Active = false
		user.Password_Hash = string(hash)
		user.Activation_Code = Common.GetRandomCode(100000, 999999)
		handler.IUserRepository.Update(user, user.Id)
	}
	handler.SmsService.SendVerificationCode("0"+strconv.FormatInt(user.Phone_Number, 10), strconv.Itoa(user.Activation_Code))

	tokenExp := time.Now().Add(2 * time.Minute)
	token, _ := Common.GenerateToken(user.Id, tokenExp)

	return (Common.ResponseModel[string]{}).ApiResponse(token, Common.SUCCESS)
}
