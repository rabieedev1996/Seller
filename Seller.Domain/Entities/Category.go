package Entities

import "time"

type Category struct {
	Id                int
	Title             string
	Summary           string
	Like_Count        int
	Parent_Cat_Id     int
	Selectable_Prop   string
	Created_At        time.Time
	Jalali_Created_At int
}
