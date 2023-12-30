package database

type Product struct {
	Id           int    `gorm:"primarykey" json:"id"`
	Product_name string `gorm:"type:varchar(255)" json:"product_name"`
	Description  string `gorm:"type:varchar(255)" json:"description"`
	Quantity     int    `json:"quantity"`
}
