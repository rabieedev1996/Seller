package GetCategories

type GetCategoriesVm struct {
	Id          int
	Title       string
	Summary     string
	LikeCount   int `mapper:"Like_Count"`
	ParentCatId int `mapper:"Parent_Cat_Id"`
}
