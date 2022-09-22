package Entities

import "time"

type CategoryProperty struct {
	Id                int
	Category_Id       int
	Title             string
	Prop_Value        string
	Created_At        time.Time
	Jalali_Created_At int
}
