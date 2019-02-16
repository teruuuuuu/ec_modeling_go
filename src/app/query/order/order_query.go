package product_query

import (
	coupon_model "../../domain/coupon/model"
	"github.com/jinzhu/gorm"
)

type ItemsView struct {
	ItemId uint
	Name   string
	Price  uint
	Number uint
}

type Coupon struct {
	CouponId     uint
	CouponNumber string
	Discount     uint
}

func CartItems(db *gorm.DB, userId uint) []ItemsView {
	var ret []ItemsView
	db.Table("orders").Select("items.item_id, products.name, items.price, items.number").
		Joins("join items on orders.order_id = items.order_id").
		Joins("join products on products.product_id = items.product_id").
		Where("orders.user_id = ? and orders.order_status = 1", userId).Scan(&ret)
	return ret
}

func ValidCoupon(db *gorm.DB, couponNumber string) (bool, *coupon_model.Coupon) {
	var ret *coupon_model.Coupon = &coupon_model.Coupon{}
	db.Table("coupons").Select("coupon_id, coupon_number, discount").
		Where("coupon_number = ? and coupon_id not in (select coupon_id from used_coupons)", couponNumber).
		FirstOrInit(ret, coupon_model.Coupon{CouponId: 0})
	if ret.CouponId == 0 {
		return false, ret
	} else {
		return true, ret
	}
}
