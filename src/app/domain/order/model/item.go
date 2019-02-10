package order_model

import "time"

type Item struct {
	ItemId     uint `gorm:"primary_key"`
	OrderId    uint `gorm:"index"`
	ProductId  uint
	Price      uint
	Number     uint
	UpdateDate time.Time
}
