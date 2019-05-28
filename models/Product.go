package models

type Product struct {
	Id   string `gorm:"primary_key"`
	Name string `gorm:"not_null"`
	Quantity int `gorm:"not_null"`
	Price float64 `gorm:"not_null"`
}
