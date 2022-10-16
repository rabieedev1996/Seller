package Entities

import "time"

type File struct {
	Id                string
	File_Name         string
	File_Length       int64
	Content_Type      string
	Extension         string
	Relation_Type     int
	Relation_Id       int
	Base_Url          string
	Relative_Url      string
	Created_At        time.Time
	Jalali_Created_At int
	Is_Deleted        bool
}
