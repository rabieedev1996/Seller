package Entities

import "time"

type SelectableValue struct {
	Id                int
	Category_Id       int
	Product_Id        int
	Prop_Value        string
	Prop_Value_Title  string
	Created_At        time.Time
	Jalali_Created_At int
}
