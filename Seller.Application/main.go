package main

import (
	"Seller/Config/InjectionConfig"
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"reflect"
	"time"
)

func main() {
	user := Entities.User{
		Id:                "1111",
		First_Name:        "Test",
		Last_Name:         "Testi",
		Phone_Number:      9128003185,
		Activation_Code:   111,
		Username:          "test",
		Email:             "mrabiee1375@gmail.com",
		Created_At:        time.Now(),
		Jalali_Created_At: 14011003,
	}

	var userRepository Presistence.IUserRepository
	userRepository = InjectionConfig.PrepareObject[Presistence.IUserRepository](reflect.TypeOf(userRepository), "")
	userRepository.Create(user)
}
