package order_model

type BankPay struct {
	BankPayId   uint  `gorm:"primary_key"`
	PaymentId   *uint `gorm:"index"`
	BankAccount string
}
