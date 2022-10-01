package GetCategoryProducts

type GetCategoryProductsQuery struct {
	CategoryId int
	Start      int
	Count      int
	SortType   string
}
