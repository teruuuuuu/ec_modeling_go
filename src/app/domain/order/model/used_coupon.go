package order_model

type UsedCoupon struct {
	OrderId  *uint `gorm:"index"`
	CouponId *uint `gorm:"index"`
	Discount uint
}
