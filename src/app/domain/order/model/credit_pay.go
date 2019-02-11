package order_model

type CreditPay struct {
	CreditPayId uint  `gorm:"primary_key"`
	PaymentId   *uint `gorm:"index"`
}
