package Entities

import "time"

type Warehouse struct {
	Id                 int
	Product_Id         int
	Selectable_Prop_Id int
	Price              int
	Exist_Count        int
	Created_At         time.Time
	Jalali_Created_At  int
}
