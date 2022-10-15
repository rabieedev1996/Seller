package GetCategoryProducts

type GetCategoryProductsVm struct {
	Id         int
	CategoryId int `mapper:"Category_Id"`
	Title      string
	Summary    string
	MinPrice   int64 `mapper:"Min_Price"`
	MaxPrice   int64 `mapper:"Max_Price"`
	LikeCount  int   `mapper:"Like_Count"`
	RankNumber float32
}
