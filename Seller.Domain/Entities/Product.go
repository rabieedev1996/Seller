package Entities

import "time"

type Product struct {
	Id                int
	Category_Id       int
	Title             string
	Summary           string
	Min_Price         int64
	Max_Price         int64
	Like_Count        int
	RankNumber        float32
	Created_At        time.Time
	Jalali_Created_At int
	Is_Deleted        bool
}
