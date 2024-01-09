package model

// using `form:"name" is a must`
type Data struct {
	ID          int    `form:"id" gorm:"primarykey"`
	Name        string `form:"name"`
	Produsen    string `form:"produsen"`
	Description string `form:"description"`
	Quantity    int    `form:"quantity"`
	Area        string `form:"area"`
}
