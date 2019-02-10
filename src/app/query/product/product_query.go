package product_query

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type ProductView struct {
	ProductId   uint
	Name        string
	Price       uint
	Description string
}

func Search(db *gorm.DB, name string) []ProductView {
	fmt.Println("search start")
	var ret []ProductView
	db.Table("products").Select("products.product_id, products.name, products.price, product_infos.description").
		Joins("left join product_infos on products.product_id = product_infos.product_id").
		Where("products.name LIKE ?", "%"+name+"%").Scan(&ret)
	return ret
}
