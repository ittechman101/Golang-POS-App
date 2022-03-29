package pos

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductId  int64   `gorm:"Not Null" json:"productid"`
	Name       string  `json:"name"`
	Stock      int64   `json:"stock"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	Category   string  `json:"category"`
	CategoryId int64   `json:"categoryId"`
	Discount   string  `json:"discount"`
}
