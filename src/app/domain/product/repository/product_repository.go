package product_repository

import (
	"github.com/jinzhu/gorm"

	product_aggregate "../aggregate"
	product_model "../model"
)

func DeleteAll(db *gorm.DB) {
	db.Delete(product_model.ProductInfo{})
	db.Delete(product_model.Product{})
}

func Save(db *gorm.DB, pa *product_aggregate.ProductAggregate) {
	db.Save(&pa.Product)
	pa.ProductInfo.ProductId = &pa.Product.ProductId
	db.Save(&pa.ProductInfo)
}
