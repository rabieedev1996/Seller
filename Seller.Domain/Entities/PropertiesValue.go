package Entities

import "time"

type PropertiesValue struct {
	Id                int
	Category_Id       int
	Product_Id        int
	Prop_Value        string
	Created_At        time.Time
	Jalali_Created_At int
}
