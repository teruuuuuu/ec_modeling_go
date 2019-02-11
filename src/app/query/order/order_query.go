package product_query

import (
	"github.com/jinzhu/gorm"
)

type ItemsView struct {
	ItemId uint
	Name   string
	Price  uint
	Number uint
}

func CartItems(db *gorm.DB, userId uint) []ItemsView {
	var ret []ItemsView
	db.Table("orders").Select("items.item_id, products.name, items.price, items.number").
		Joins("join items on orders.order_id = items.order_id").
		Joins("join products on products.product_id = items.product_id").
		Where("orders.user_id = ? and orders.order_status = 1", userId).Scan(&ret)
	return ret
}
