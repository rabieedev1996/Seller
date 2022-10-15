package GetProductDetail

type GetProductQueryVM struct {
	Id          int
	CategoryId  int `mapper:"Category_Id"`
	Title       string
	Summary     string
	MinPrice    int64   `mapper:"Min_Price"`
	MaxPrice    int64   `mapper:"Max_Price"`
	LikeCount   int     `mapper:"Like_Count"`
	RankNumber  float32 `mapper:"Rank_Number"`
	ExistCount  int
	Properties  []ProductPropertiesVM
	Selectables ProductSelectablesVM
}

type ProductPropertiesVM struct {
	Title         string
	PropertyValue string `mapper:"Prop_Value"`
}

type ProductSelectablesVM struct {
	Title  string
	Values []ProductSelectablesValuesVM
}

type ProductSelectablesValuesVM struct {
	ValueTitle string
	ValueKey   string
	ExistCount int
}
