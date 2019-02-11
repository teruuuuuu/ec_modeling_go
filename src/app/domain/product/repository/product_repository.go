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

type ProductView struct {
	ProductId   uint
	Name        string
	Price       uint
	Description string
}

func FindById(db *gorm.DB, productId uint) (bool, *product_aggregate.ProductAggregate) {
	var ret *product_aggregate.ProductAggregate = &product_aggregate.ProductAggregate{}
	var product *product_model.Product = &product_model.Product{}
	var productInfo *product_model.ProductInfo = &product_model.ProductInfo{}
	db.Where("product_id = ?", productId).FirstOrInit(&product, product_model.Product{ProductId: 0})
	db.Where("product_id = ?", productId).First(&productInfo, product_model.ProductInfo{ProductId: nil})
	if product.ProductId == 0 || productInfo.ProductId == nil {
		return false, ret
	}
	ret.Product = product
	ret.ProductInfo = productInfo
	return true, ret
}

func Save(db *gorm.DB, pa *product_aggregate.ProductAggregate) {
	db.Save(&pa.Product)
	pa.ProductInfo.ProductId = &pa.Product.ProductId
	db.Save(&pa.ProductInfo)
}
