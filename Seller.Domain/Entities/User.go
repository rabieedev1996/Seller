package Entities

import "time"

type User struct {
	Id                string
	First_Name        string
	Last_Name         string
	Phone_Number      int64
	Activation_Code   int
	Username          string
	Password_Hash     string
	Email             string
	Created_At        time.Time
	Jalali_Created_At int
}
