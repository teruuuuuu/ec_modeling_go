package coupon_repository

import (
	"github.com/jinzhu/gorm"

	coupon_model "../model"
)

func DeleteAll(db *gorm.DB) {
	db.Delete(coupon_model.Coupon{})
}

func Save(db *gorm.DB, c *coupon_model.Coupon) {
	db.Save(&c)
}
