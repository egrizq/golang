package model

type User struct {
	ID   int `gorm:"primaryKey"`
	User string
	News []News `gorm:"foreignKey:UserID"`
}

type News struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Main   string
	UserID int
}
