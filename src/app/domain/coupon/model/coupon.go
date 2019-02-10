package coupon_model

type Coupon struct {
	CouponId     uint `gorm:"primary_key"`
	CouponNumber string
	Discount     uint
}
